package filters

func (f *Filters) Merge(other Filters) {
	f.FilterContainer.Merge(other.FilterContainer)
}

func (fc *FilterContainer) Merge(other FilterContainer) {
	fc.FilterRefs = append(fc.FilterRefs, other.FilterRefs...)

	fc.Always = append(fc.Always, other.Always...)
	fc.Never = append(fc.Never, other.Never...)
	fc.Time = append(fc.Time, other.Time...)
	fc.MatchStarted = append(fc.MatchStarted, other.MatchStarted...)
	fc.MatchRunning = append(fc.MatchRunning, other.MatchRunning...)
	fc.MatchFinished = append(fc.MatchFinished, other.MatchFinished...)
	fc.Completed = append(fc.Completed, other.Completed...)
	fc.Captured = append(fc.Captured, other.Captured...)

	fc.FlagCarried = append(fc.FlagCarried, other.FlagCarried...)
	fc.FlagDropped = append(fc.FlagDropped, other.FlagDropped...)
	fc.FlagReturned = append(fc.FlagReturned, other.FlagReturned...)
	fc.FlagCaptured = append(fc.FlagCaptured, other.FlagCaptured...)

	fc.Void = append(fc.Void, other.Void...)
	fc.Blocks = append(fc.Blocks, other.Blocks...)

	fc.Material = append(fc.Material, other.Material...)
	fc.StructuralLoad = append(fc.StructuralLoad, other.StructuralLoad...)

	fc.Spawn = append(fc.Spawn, other.Spawn...)
	fc.Mob = append(fc.Mob, other.Mob...)
	fc.Entity = append(fc.Entity, other.Entity...)

	fc.CarryingFlag = append(fc.CarryingFlag, other.CarryingFlag...)
	fc.Score = append(fc.Score, other.Score...)
	fc.Rank = append(fc.Rank, other.Rank...)

	fc.Participating = append(fc.Participating, other.Participating...)
	fc.Observing = append(fc.Observing, other.Observing...)
	fc.KillStreak = append(fc.KillStreak, other.KillStreak...)
	fc.Lives = append(fc.Lives, other.Lives...)
	fc.Crouching = append(fc.Crouching, other.Crouching...)
	fc.Walking = append(fc.Walking, other.Walking...)
	fc.Sprinting = append(fc.Sprinting, other.Sprinting...)
	fc.Grounded = append(fc.Grounded, other.Grounded...)
	fc.Flying = append(fc.Flying, other.Flying...)
	fc.Alive = append(fc.Alive, other.Alive...)
	fc.Dead = append(fc.Dead, other.Dead...)
	fc.CanFly = append(fc.CanFly, other.CanFly...)
	fc.Team = append(fc.Team, other.Team...)
	fc.Class = append(fc.Class, other.Class...)
	fc.Effect = append(fc.Effect, other.Effect...)

	fc.Carrying = append(fc.Carrying, other.Carrying...)
	fc.Holding = append(fc.Holding, other.Holding...)
	fc.Wearing = append(fc.Wearing, other.Wearing...)

	fc.Cause = append(fc.Cause, other.Cause...)
	fc.Random = append(fc.Random, other.Random...)

	fc.Relation = append(fc.Relation, other.Relation...)

	fc.Countdown = append(fc.Countdown, other.Countdown...)
	fc.After = append(fc.After, other.After...)
	fc.Pulse = append(fc.Pulse, other.Pulse...)

	fc.Variable = append(fc.Variable, other.Variable...)
	fc.Offset = append(fc.Offset, other.Offset...)

	fc.Players = append(fc.Players, other.Players...)

	fc.Not = append(fc.Not, other.Not...)
	fc.One = append(fc.One, other.One...)
	fc.All = append(fc.All, other.All...)
	fc.Any = append(fc.Any, other.Any...)
	fc.Allow = append(fc.Allow, other.Allow...)
	fc.Deny = append(fc.Deny, other.Deny...)
	fc.SameTeam = append(fc.SameTeam, other.SameTeam...)
	fc.Victim = append(fc.Victim, other.Victim...)
	fc.Attacker = append(fc.Attacker, other.Attacker...)
	fc.Player = append(fc.Player, other.Player...)
}
