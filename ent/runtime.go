// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/lukepeltier/discordle/ent/blacklist"
	"github.com/lukepeltier/discordle/ent/discordmember"
	"github.com/lukepeltier/discordle/ent/message"
	"github.com/lukepeltier/discordle/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	blacklistFields := schema.Blacklist{}.Fields()
	_ = blacklistFields
	// blacklistDescBad is the schema descriptor for bad field.
	blacklistDescBad := blacklistFields[0].Descriptor()
	// blacklist.BadValidator is a validator for the "bad" field. It is called by the builders before save.
	blacklist.BadValidator = blacklistDescBad.Validators[0].(func(string) error)
	discordmemberFields := schema.DiscordMember{}.Fields()
	_ = discordmemberFields
	// discordmemberDescUsername is the schema descriptor for username field.
	discordmemberDescUsername := discordmemberFields[0].Descriptor()
	// discordmember.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	discordmember.UsernameValidator = discordmemberDescUsername.Validators[0].(func(string) error)
	// discordmemberDescBlacklisted is the schema descriptor for blacklisted field.
	discordmemberDescBlacklisted := discordmemberFields[3].Descriptor()
	// discordmember.DefaultBlacklisted holds the default value on creation for the blacklisted field.
	discordmember.DefaultBlacklisted = discordmemberDescBlacklisted.Default.(bool)
	messageFields := schema.Message{}.Fields()
	_ = messageFields
	// messageDescText is the schema descriptor for text field.
	messageDescText := messageFields[0].Descriptor()
	// message.TextValidator is a validator for the "text" field. It is called by the builders before save.
	message.TextValidator = messageDescText.Validators[0].(func(string) error)
	// messageDescDateSent is the schema descriptor for date_sent field.
	messageDescDateSent := messageFields[1].Descriptor()
	// message.DateSentValidator is a validator for the "date_sent" field. It is called by the builders before save.
	message.DateSentValidator = messageDescDateSent.Validators[0].(func(string) error)
}
