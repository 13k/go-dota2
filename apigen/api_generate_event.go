package main

import (
	"fmt"

	pb "github.com/13k/go-steam-resources/protobuf/dota2"
)

type generatedEventHandler struct {
	msgID     pb.EDOTAGCMsg
	eventName string
	eventType *ProtoType
}

// buildGeneratedEventHandler builds a generated event handler.
func buildGeneratedEventHandler(
	msgID pb.EDOTAGCMsg,
	protoMap map[string]*ProtoType,
	eventImports map[string]struct{},
) (*generatedEventHandler, error) {
	eventProtoType, err := LookupMessageProtoType(protoMap, msgID)
	if err != nil {
		return nil, err
	}
	eventImports[eventProtoType.Pak.Path()] = struct{}{}

	return &generatedEventHandler{
		msgID:     msgID,
		eventName: GetMessageEventName(msgID),
		eventType: eventProtoType,
	}, nil
}

func (g *generatedEventHandler) generateComment() string {
	return fmt.Sprintf("// %s event.\n// MessageID: %s\n", g.eventName, g.msgID.String())
}
