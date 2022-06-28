// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ChannelsColumns holds the columns for the "channels" table.
	ChannelsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "display_name", Type: field.TypeString, Unique: true},
		{Name: "image_path", Type: field.TypeString},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
	}
	// ChannelsTable holds the schema information for the "channels" table.
	ChannelsTable = &schema.Table{
		Name:       "channels",
		Columns:    ChannelsColumns,
		PrimaryKey: []*schema.Column{ChannelsColumns[0]},
	}
	// QueuesColumns holds the columns for the "queues" table.
	QueuesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "live_archive", Type: field.TypeBool, Default: false},
		{Name: "on_hold", Type: field.TypeBool, Default: false},
		{Name: "video_processing", Type: field.TypeBool, Default: true},
		{Name: "chat_processing", Type: field.TypeBool, Default: true},
		{Name: "processing", Type: field.TypeBool, Default: true},
		{Name: "task_vod_create_folder", Type: field.TypeEnum, Nullable: true, Enums: []string{"success", "processing", "waiting", "Error"}, Default: "waiting"},
		{Name: "task_vod_download_thumbnail", Type: field.TypeEnum, Nullable: true, Enums: []string{"success", "processing", "waiting", "Error"}, Default: "waiting"},
		{Name: "task_vod_save_info", Type: field.TypeEnum, Nullable: true, Enums: []string{"success", "processing", "waiting", "Error"}, Default: "waiting"},
		{Name: "task_video_download", Type: field.TypeEnum, Nullable: true, Enums: []string{"success", "processing", "waiting", "Error"}, Default: "waiting"},
		{Name: "task_video_move", Type: field.TypeEnum, Nullable: true, Enums: []string{"success", "processing", "waiting", "Error"}, Default: "waiting"},
		{Name: "task_chat_download", Type: field.TypeEnum, Nullable: true, Enums: []string{"success", "processing", "waiting", "Error"}, Default: "waiting"},
		{Name: "task_chat_render", Type: field.TypeEnum, Nullable: true, Enums: []string{"success", "processing", "waiting", "Error"}, Default: "waiting"},
		{Name: "task_chat_move", Type: field.TypeEnum, Nullable: true, Enums: []string{"success", "processing", "waiting", "Error"}, Default: "waiting"},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "vod_queue", Type: field.TypeUUID, Unique: true},
	}
	// QueuesTable holds the schema information for the "queues" table.
	QueuesTable = &schema.Table{
		Name:       "queues",
		Columns:    QueuesColumns,
		PrimaryKey: []*schema.Column{QueuesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "queues_vods_queue",
				Columns:    []*schema.Column{QueuesColumns[16]},
				RefColumns: []*schema.Column{VodsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "role", Type: field.TypeEnum, Enums: []string{"admin", "editor", "archiver", "user"}, Default: "user"},
		{Name: "webhook", Type: field.TypeString, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// VodsColumns holds the columns for the "vods" table.
	VodsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "ext_id", Type: field.TypeString},
		{Name: "platform", Type: field.TypeEnum, Enums: []string{"twitch", "youtube"}, Default: "twitch"},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"archive", "live", "highlight", "upload", "clip"}, Default: "archive"},
		{Name: "title", Type: field.TypeString},
		{Name: "duration", Type: field.TypeInt, Default: 0},
		{Name: "views", Type: field.TypeInt, Default: 0},
		{Name: "resolution", Type: field.TypeString, Nullable: true},
		{Name: "processing", Type: field.TypeBool, Default: false},
		{Name: "thumbnail_path", Type: field.TypeString, Nullable: true},
		{Name: "web_thumbnail_path", Type: field.TypeString},
		{Name: "video_path", Type: field.TypeString},
		{Name: "chat_path", Type: field.TypeString, Nullable: true},
		{Name: "chat_video_path", Type: field.TypeString, Nullable: true},
		{Name: "info_path", Type: field.TypeString, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "channel_vods", Type: field.TypeUUID},
	}
	// VodsTable holds the schema information for the "vods" table.
	VodsTable = &schema.Table{
		Name:       "vods",
		Columns:    VodsColumns,
		PrimaryKey: []*schema.Column{VodsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "vods_channels_vods",
				Columns:    []*schema.Column{VodsColumns[17]},
				RefColumns: []*schema.Column{ChannelsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ChannelsTable,
		QueuesTable,
		UsersTable,
		VodsTable,
	}
)

func init() {
	QueuesTable.ForeignKeys[0].RefTable = VodsTable
	VodsTable.ForeignKeys[0].RefTable = ChannelsTable
}
