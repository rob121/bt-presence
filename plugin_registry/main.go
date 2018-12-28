package plugin_registry

type NotificationPlugin interface {
    Arrived(mac string,room string) 
    Departed(mac string,room string) 
}

var Notifiers = []NotificationPlugin{}

func RegisterNotifcation(n NotificationPlugin) {
    Notifiers = append(Notifiers, n)
}

