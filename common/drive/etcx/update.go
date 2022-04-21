package etcx

import "context"

func Update(info *EtcEnv, service, val string) error {
	switch info.Type {
	case "etcd":
		if client, err := NewEtcd(info); err != nil {
			return err
		} else {
			_, err = client.Client.KV.Put(context.TODO(), info.Prefix+service, val)
			return err
		}
	}
	return nil
}
