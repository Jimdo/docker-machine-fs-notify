package main // import "github.com/Jimdo/docker-machine-fs-notify"
import (
	"errors"
	"os"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/docker/machine/commands/mcndirs"
	"github.com/docker/machine/libmachine"
	"github.com/docker/machine/libmachine/state"
	"github.com/howeyc/fsnotify"
	"gopkg.in/alecthomas/kingpin.v2"
)

const (
	touchTimeFormat string = "200601021504.05" // YYYYMMDDhhmm.SS
)

type FsEvent struct {
	File    string
	ModTime time.Time
}

type DockerMachineFsNotify struct {
	DockerMachineName string
	RecentEvents      map[string]FsEvent
}

func NewDockerMachineFsNotify(dockerMachineName string) *DockerMachineFsNotify {
	return &DockerMachineFsNotify{
		DockerMachineName: dockerMachineName,
		RecentEvents:      make(map[string]FsEvent),
	}
}

func (d *DockerMachineFsNotify) NotifyVm(event FsEvent) error {
	api := libmachine.NewClient(mcndirs.GetBaseDir(), mcndirs.GetMachineCertDir())

	host, err := api.Load(d.DockerMachineName)
	if err != nil {
		return err
	}

	currentState, err := host.Driver.GetState()
	if err != nil {
		return err
	}

	if currentState != state.Running {
		return errors.New("Docker Machine " + host.Name + " is not running")
	}

	client, err := host.CreateSSHClient()
	if err != nil {
		return err
	}

	return client.Shell("touch -t " + event.ModTime.UTC().Format(touchTimeFormat) + " -m -c " + event.File)
}

func (p *DockerMachineFsNotify) ProcessEvent(fileEvent *fsnotify.FileEvent) {

	if fileEvent.IsDelete() || fileEvent.IsRename() {
		// We cannot handle delete for the moment
		return
	}

	event := FsEvent{
		File: fileEvent.Name,
	}

	info, err := os.Stat(event.File)
	if err != nil {
		log.WithFields(log.Fields{"error": err, "file": event.File}).Warn("Error reading file")
		return
	}
	event.ModTime = info.ModTime()

	existingEvent, ok := p.RecentEvents[event.File]
	if ok && existingEvent.ModTime.Equal(event.ModTime) {
		// Recursion event -> Do nothing
	} else {
		log.WithFields(log.Fields{"file": event.File}).Info("New filesystem event")
		p.RecentEvents[event.File] = event
		err = p.NotifyVm(event)
		if err != nil {
			log.WithFields(log.Fields{"error": err}).Warn("Error when notifying VM")
			return
		}
	}
}

func (p *DockerMachineFsNotify) CleanupRecentEvents(ttl time.Duration) {
	for key, event := range p.RecentEvents {
		if event.ModTime.Before(time.Now().Add(-ttl)) {
			delete(p.RecentEvents, key)
		}
	}
}

func main() {
	var (
		directory         = kingpin.Arg("directory", "").Required().String()
		dockerMachineName = kingpin.Arg("docker-machine-name", "").Required().String()
	)

	kingpin.UsageTemplate(kingpin.CompactUsageTemplate)
	kingpin.CommandLine.Help = "Forward file system events to a docker machine VM"
	kingpin.Parse()

	p := NewDockerMachineFsNotify(*dockerMachineName)

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan bool)

	// Process events
	go func() {
		for {
			select {
			case ev := <-watcher.Event:
				p.ProcessEvent(ev)

			case err := <-watcher.Error:
				log.WithFields(log.Fields{"error": err}).Warn("Error in fsnotify watcher")

			case <-time.After(1 * time.Minute):
				p.CleanupRecentEvents(10 * time.Second)
			}
		}
	}()

	err = watcher.Watch(*directory)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Fatal("Error in fsnotify watcher")
	}

	// Hang so program doesn't exit
	<-done

	watcher.Close()
}
