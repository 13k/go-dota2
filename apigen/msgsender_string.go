// Code generated by "stringer -type=MsgSender"; DO NOT EDIT.

package main

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[MsgSenderNone-0]
	_ = x[MsgSenderGC-1]
	_ = x[MsgSenderClient-2]
}

const _MsgSender_name = "MsgSenderNoneMsgSenderGCMsgSenderClient"

var _MsgSender_index = [...]uint8{0, 13, 24, 39}

func (i MsgSender) String() string {
	if i >= MsgSender(len(_MsgSender_index)-1) {
		return "MsgSender(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _MsgSender_name[_MsgSender_index[i]:_MsgSender_index[i+1]]
}
