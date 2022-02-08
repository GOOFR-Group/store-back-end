ALTER TABLE "game" 
ADD CONSTRAINT "game_state_check" 
CHECK (
	"state" = 'active'
	OR "state" = 'inactive'
	OR "state" = 'upcoming'
);