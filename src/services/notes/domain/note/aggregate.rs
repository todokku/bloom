use serde::{Serialize, Deserialize};
use diesel::{Queryable};
use crate::{
    db::schema::notes_notes,
};


#[derive(AsChangeset, Clone, Debug, Deserialize, Identifiable, Insertable, Queryable, Serialize)]
#[table_name = "notes_notes"]
#[changeset_options(treat_none_as_null = "true")]
pub struct Note {
    pub id: uuid::Uuid,
    pub created_at: chrono::DateTime<chrono::Utc>,
    pub updated_at: chrono::DateTime<chrono::Utc>,
    pub deleted_at: Option<chrono::DateTime<chrono::Utc>>,
    pub version: i64,

    pub archived_at: Option<chrono::DateTime<chrono::Utc>>,
    pub body: String,
    pub removed_at: Option<chrono::DateTime<chrono::Utc>>,
    pub title: String,

    pub owner_id: uuid::Uuid,
}


impl Note {
    // create a new, unitialized note
    pub fn new() -> Self {
        let now = chrono::Utc::now();
        return Note{
            id: uuid::Uuid::new_v4(),
            created_at: now,
            updated_at: now,
            deleted_at: None,
            version: 0,

            archived_at: None,
            body: String::new(),
            removed_at: None,
            title: String::new(),

            owner_id: uuid::Uuid::new_v4(),
        };
    }
}

impl eventsourcing::Aggregate for Note {
    fn increment_version(&mut self) {
        self.version += 1;
    }

    fn update_updated_at(&mut self, timestamp: chrono::DateTime<chrono::Utc>) {
        self.updated_at = timestamp;
    }
}
