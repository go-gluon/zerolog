// Code generated by gluon; DO NOT EDIT.
package zerolog

import config "github.com/go-gluon/gluon/config"

func (item *ZerologConfig) ReadFromMapNode(node config.MapNode) error {
	item.Debug = node.Bool("debug", item.Debug)
	item.Json = node.Bool("json", item.Json)
	return nil
}