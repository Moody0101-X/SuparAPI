package models;


type Notification struct {
	Id_       int      `json:"id"`
	Text      string   `json:"text"`
	Type      int      `json:"type"` 
	Uuid      int      `json:"uuid"`
	Actorid   int      `json:"actorid"`
	Seen      int      `json:"seen"`
	Post_id   int      `json:"post_id"`
	Link      string   `json:"link"`
	User_     AUser    `json:"user"`
	/*
	    CREATE TABLE NOTIFICATIONS (
        ID INTEGER PRIMARY KEY AUTOINCREMENT,
        Text TEXT,
        TYPE INTEGER,
        UUID INTEGER,
        ACTORID INTEGER,
        Seen INTEGER,
        Post_id INTEGER,
        Link TEXT
    )

	*/
}

func NewNot(text string, t int, uuid int, actorid int) Notification {
	var new Notification;
	
	new.Text = text;
	new.Type = t;
	new.Uuid = uuid;
	new.Actorid = actorid;
	new.Seen = 0;

	return new;
}

