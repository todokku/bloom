use crate::{domain::playlist, validators};
use diesel::{
    r2d2::{ConnectionManager, PooledConnection},
    PgConnection,
};
use eventsourcing::{Event, EventTs};
use kernel::KernelError;

#[derive(Clone, Debug)]
pub struct Create {
    pub name: String,
    pub owner_id: uuid::Uuid,
}

impl eventsourcing::Command for Create {
    type Aggregate = playlist::Playlist;
    type Event = Created;
    type Context = PooledConnection<ConnectionManager<PgConnection>>;
    type Error = KernelError;

    fn validate(
        &self,
        _ctx: &Self::Context,
        _aggregate: &Self::Aggregate,
    ) -> Result<(), Self::Error> {
        validators::playlist_name(&self.name)?;

        return Ok(());
    }

    fn build_event(
        &self,
        _ctx: &Self::Context,
        _aggregate: &Self::Aggregate,
    ) -> Result<Self::Event, Self::Error> {
        return Ok(Created {
            id: uuid::Uuid::new_v4(),
            timestamp: chrono::Utc::now(),
            name: self.name.clone(),
            owner_id: self.owner_id,
        });
    }
}

// Event
#[derive(Clone, Debug, EventTs)]
pub struct Created {
    pub timestamp: chrono::DateTime<chrono::Utc>,
    pub id: uuid::Uuid,
    pub name: String,
    pub owner_id: uuid::Uuid,
}

impl Event for Created {
    type Aggregate = playlist::Playlist;

    fn apply(&self, aggregate: Self::Aggregate) -> Self::Aggregate {
        return Self::Aggregate {
            id: self.id,
            created_at: self.timestamp,
            updated_at: self.timestamp,
            deleted_at: None,
            version: 0,
            name: self.name.clone(),
            owner_id: self.owner_id,
        };
    }
}
