package libeffone

type PacketHeader struct {
	PacketFormat            uint16  // 2020
	GameMajorVersion        uint8   // Game major version - "X.00"
	GameMinorVersion        uint8   // Game minor version - "1.XX"
	PacketVersion           uint8   // Version of this packet type, all start from 1
	PacketID                uint8   // Identifier for the packet type, see below
	SessionUID              uint64  // Unique identifier for the session
	SessionTime             float32 // Session timestamp
	FrameIdentifier         uint32  // Identifier for the frame the data was retrieved on
	PlayerCarIndex          uint8   // Index of player's car in the array
	SecondaryPlayerCarIndex uint8   // Index of secondary player's car in the array (splitscreen) // 255 if no second player
}

type CarMotionData struct {
	WorldPositionX         float32 // World space X position
	WorldPositionY         float32 // World space Y position
	WorldPositionZ         float32 // World space Z position
	WorldVelocityX         float32 // Velocity in world space X
	WorldVelocityY         float32 // Velocity in world space Y
	WorldVelocityZ         float32 // Velocity in world space Z
	WorldForwardDirectionX int16   // World space forward X direction (normalised)
	WorldForwardDirectionY int16   // World space forward Y direction (normalised)
	WorldForwardDirectionZ int16   // World space forward Z direction (normalised)
	WorldRightDirectionX   int16   // World space right X direction (normalised)
	WorldRightDirectionY   int16   // World space right Y direction (normalised)
	WorldRightDirectionZ   int16   // World space right Z direction (normalised)
	GForceLateral          float32 // Lateral G-Force component
	GForceLongitudinal     float32 // Longitudinal G-Force component
	GForceVertical         float32 // Vertical G-Force component
	Yaw                    float32 // Yaw angle in radians
	Pitch                  float32 // Pitch angle in radians
	Roll                   float32 // Roll angle in radians
}

type PacketMotionData struct {
	Header                 PacketHeader      // Header
	CarMotionData          [22]CarMotionData // Data for all cars on track // Extra player car ONLY data below:
	SuspensionPosition     [4]float32        // Note: All wheel arrays have the following order:
	SuspensionVelocity     [4]float32        // RL, RR, FL, FR
	SuspensionAcceleration [4]float32        // RL, RR, FL, FR
	WheelSpeed             [4]float32        // Speed of each wheel
	WheelSlip              [4]float32        // Slip ratio for each wheel
	LocalVelocityX         float32           // Velocity in local space
	LocalVelocityY         float32           // Velocity in local space
	LocalVelocityZ         float32           // Velocity in local space
	AngularVelocityX       float32           // Angular velocity x-component
	AngularVelocityY       float32           // Angular velocity y-component
	AngularVelocityZ       float32           // Angular velocity z-component
	AngularAccelerationX   float32           // Angular velocity x-component
	AngularAccelerationY   float32           // Angular velocity y-component
	AngularAccelerationZ   float32           // Angular velocity z-component
	FrontWheelsAngle       float32           // Current front wheels angle in radians
}

type MarshalZone struct {
	ZoneStart float32 // Fraction (0..1) of way through the lap the marshal zone starts
	ZoneFlag  int8    // -1 = invalid/unknown, 0 = none, 1 = green, 2 = blue, 3 = yellow, 4 = red
}

type WeatherForecastSample struct {
	SessionType      uint8 // 0 = unknown, 1 = P1, 2 = P2, 3 = P3, 4 = Short P, 5 = Q1 // 6 = Q2, 7 = Q3, 8 = Short Q, 9 = OSQ, 10 = R, 11 = R2 // 12 = Time Trial
	TimeOffset       uint8 // Time in minutes the forecast is for
	Weather          uint8 // Weather - 0 = clear, 1 = light cloud, 2 = overcast // 3 = light rain, 4 = heavy rain, 5 = storm
	TrackTemperature int8  // Track temp. in degrees celsius
	AirTemperature   int8  // Air temp. in degrees celsius
}

type PacketSessionData struct {
	Header                    PacketHeader              // Header
	Weather                   uint8                     // Weather - 0 = clear, 1 = light cloud, 2 = overcast // 3 = light rain, 4 = heavy rain, 5 = storm
	TrackTemperature          int8                      // Track temp. in degrees celsius
	AirTemperature            int8                      // Air temp. in degrees celsius
	TotalLaps                 uint8                     // Total number of laps in this race
	TrackLength               uint16                    // Track length in metres
	SessionType               uint8                     // 0 = unknown, 1 = P1, 2 = P2, 3 = P3, 4 = Short P // 5 = Q1, 6 = Q2, 7 = Q3, 8 = Short Q, 9 = OSQ // 10 = R, 11 = R2, 12 = Time Trial
	TrackID                   int8                      // -1 for unknown, 0-21 for tracks, see appendix
	Formula                   uint8                     // Formula, 0 = F1 Modern, 1 = F1 Classic, 2 = F2, // 3 = F1 Generic
	SessionTimeLeft           uint16                    // Time left in session in seconds
	SessionDuration           uint16                    // Session duration in seconds
	PitSpeedLimit             uint8                     // Pit speed limit in kilometres per hour
	GamePaused                uint8                     // Whether the game is paused
	IsSpectating              uint8                     // Whether the player is spectating
	SpectatorCarIndex         uint8                     // Index of the car being spectated
	SLIProNativeSupport       uint8                     // SLI Pro support, 0 = inactive, 1 = active
	NumMarshalZones           uint8                     // Number of marshal zones to follow
	MarshalZone               [21]MarshalZone           // List of marshal zones – max 21
	SafetyCarStatus           uint8                     // 0 = no safety car, 1 = full safety car // 2 = virtual safety car
	NetworkGame               uint8                     // 0 = offline, 1 = online
	NumWeatherForecastSamples uint8                     // Number of weather samples to follow
	WeatherForecastSamples    [20]WeatherForecastSample // Array of weather forecast samples
}

type LapData struct {
	LastLapTime                float32 // Last lap time in seconds
	CurrentLapTime             float32 // Current time around the lap in seconds
	Sector1TimeInMs            uint16  // Sector 1 time in milliseconds
	Sector2TimeInMs            uint16  // Sector 2 time in milliseconds
	BestLapTime                float32 // Best lap time of the session in seconds
	BestLapNum                 uint8   // Lap number best time achieved on
	BestLapSector1TimeInMs     uint16  // Sector 1 time of best lap in the session in milliseconds
	BestLapSector2TimeInMs     uint16  // Sector 2 time of best lap in the session in milliseconds
	BestLapSector3TimeInMs     uint16  // Sector 3 time of best lap in the session in milliseconds
	BestOverallSector1TimeInMs uint16  // Best overall sector 1 time of the session in milliseconds
	BestOverallSector1LapNum   uint8   // Lap number best overall sector 1 time achieved on
	BestOverallSector2TimeInMs uint16  // Best overall sector 2 time of the session in milliseconds
	BestOverallSector2LapNum   uint8   // Lap number best overall sector 2 time achieved on
	BestOverallSector3TimeInMs uint16  // Best overall sector 3 time of the session in milliseconds
	BestOverallSector3LapNum   uint8   // Lap number best overall sector 3 time achieved on

	LapDistance       float32 // Distance vehicle is around current lap in metres – could // be negative if line hasn’t been crossed yet
	TotalDistance     float32 // Total distance travelled in session in metres – could // be negative if line hasn’t been crossed yet
	SafetyCarDelta    float32 // Delta in seconds for safety car
	CarPosition       uint8   // Car race position
	CurrentLapNum     uint8   // Current lap number
	PitStatus         uint8   // 0 = none, 1 = pitting, 2 = in pit area
	Sector            uint8   // 0 = sector1, 1 = sector2, 2 = sector3
	CurrentLapInvalid uint8   // Current lap invalid - 0 = valid, 1 = invalid
	Penalties         uint8   // Accumulated time penalties in seconds to be added
	GridPosition      uint8   // Grid position the vehicle started the race in
	DriverStatus      uint8   // Status of driver - 0 = in garage, 1 = flying lap // 2 = in lap, 3 = out lap, 4 = on track
	ResultStatus      uint8   // Result status - 0 = invalid, 1 = inactive, 2 = active // 3 = finished, 4 = disqualified, 5 = not classified // 6 = retired
}

type PacketLapData struct {
	Header  PacketHeader // Header
	LapData [22]LapData  // Lap data for all cars on track
}

type EventDataDetails struct {
	FastestLap struct {
		VehicleIndex uint8   // Vehicle index of car achieving fastest lap
		LapTime      float32 // Lap time is in seconds
	}

	Retirement struct {
		VehicleIndex uint8 // Vehicle index of car retiring
	}

	TeamMateInPits struct {
		VehicleIndex uint8 // Vehicle index of team mate
	}

	RaceWinner struct {
		VehicleIndex uint8 // Vehicle index of the race winner
	}

	Penalty struct {
		PenaltyType       uint8 // Penalty type – see Appendices
		InfringementType  uint8 // Infringement type – see Appendices
		VehicleIndex      uint8 // Vehicle index of the car the penalty is applied to
		OtherVehicleIndex uint8 // Vehicle index of the other car involved
		Time              uint8 // Time gained, or time spent doing action in seconds
		LapNum            uint8 // Lap the penalty occurred on
		PlacesGained      uint8 // Number of places gained by this
	}

	SpeedTrap struct {
		VehicleIndex uint8   // Vehicle index of the vehicle triggering speed trap
		Speed        float32 // Top speed achieved in kilometres per hour
	}
}

type PacketEventData struct {
	Header          PacketHeader     // Header
	EventStringCode [4]uint8         // Event string code, see below
	EventDetails    EventDataDetails // Event details - should be interpreted differently // for each type
}

type ParticipantData struct {
	AIControlled  uint8     // Whether the vehicle is AI (1) or Human (0) controlled
	DriverID      uint8     // Driver id - see appendix
	TeamID        uint8     // Team id - see appendix
	RaceNumber    uint8     // Race number of the car
	Nationality   uint8     // Nationality of the driver
	Name          [48]uint8 // Name of participant in UTF-8 format – null terminated // Will be truncated with … (U+2026) if too long
	YourTelemetry uint8     // The player's UDP setting, 0 = restricted, 1 = public
}

type PacketParticipantsData struct {
	Header        PacketHeader // Header
	NumActiveCars uint8        // Number of active cars in the data – should match number of // cars on HUD
	Participants  [22]ParticipantData
}

type CarSetupData struct {
	FrontWing              uint8   // Front wing aero
	RearWing               uint8   // Rear wing aero
	OnThrottle             uint8   // Differential adjustment on throttle (percentage)
	OffThrottle            uint8   // Differential adjustment off throttle (percentage)
	FrontCamber            float32 // Front camber angle (suspension geometry)
	RearCamber             float32 // Rear camber angle (suspension geometry)
	FrontToe               float32 // Front toe angle (suspension geometry)
	RearToe                float32 // Rear toe angle (suspension geometry)
	FrontSuspension        uint8   // Front suspension
	RearSuspension         uint8   // Rear suspension
	FrontAntiRollBar       uint8   // Front anti-roll bar
	RearAntiRollBar        uint8   // Front anti-roll bar
	FrontSuspensionHeight  uint8   // Front ride height
	RearSuspensionHeight   uint8   // Rear ride height
	BrakePressure          uint8   // Brake pressure (percentage)
	BrakeBias              uint8   // Brake bias (percentage)
	RearLeftTyrePressure   float32 // Rear left tyre pressure (PSI)
	RearRightTyrePressure  float32 // Rear right tyre pressure (PSI)
	FrontLeftTyrePressure  float32 // Front left tyre pressure (PSI)
	FrontRightTyrePressure float32 // Front right tyre pressure (PSI)
	Ballast                uint8   // Ballast
	FuelLoad               float32 // Fuel load
}

type PacketCarSetupData struct {
	Header    PacketHeader // Header
	CarSetups [22]CarSetupData
}

type CarTelemetryData struct {
	Speed                   uint16     // Speed of car in kilometres per hour
	Throttle                float32    // Amount of throttle applied (0.0 to 1.0)
	Steer                   float32    // Steering (-1.0 (full lock left) to 1.0 (full lock right))
	Brake                   float32    // Amount of brake applied (0.0 to 1.0)
	Clutch                  uint8      // Amount of clutch applied (0 to 100)
	Gear                    int8       // Gear selected (1-8, N=0, R=-1)
	EngineRPM               uint16     // Engine RPM
	DRS                     uint8      // 0 = off, 1 = on
	RevLightsIndicator      uint8      // Rev lights indicator (percentage)
	BrakesTemperature       [4]uint16  // Brakes temperature (celsius)
	TyresSurfaceTemperature [4]uint8   // Tyres surface temperature (celsius)
	TyresInnerTemperature   [4]uint8   // Tyres inner temperature (celsius)
	EngineTemperature       uint16     // Engine temperature (celsius)
	TyresPressure           [4]float32 // Tyres pressure (PSI)
	SurfaceType             [4]uint8   // Driving surface, see appendices
}

type PacketCarTelemetryData struct {
	Header                       PacketHeader // Header
	CarTelemetryData             [22]CarTelemetryData
	ButtonStatus                 uint32 // Bit flags specifying which buttons are being pressed // currently - see appendices
	MFDPanelIndex                uint8  // Index of MFD panel open - 255 = MFD closed // Single player, race – 0 = Car setup, 1 = Pits // 2 = Damage, 3 =  Engine, 4 = Temperatures // May vary depending on game mode
	MFDPanelIndexSecondaryPlayer uint8  // See above
	SuggestedGear                int8   // Suggested gear for the player (1-8) // 0 if no gear suggested
}

type CarStatusData struct {
	TractionControl         uint8    // 0 (off) - 2 (high)
	AntiLockBrakes          uint8    // 0 (off) - 1 (on)
	FuelMix                 uint8    // Fuel mix - 0 = lean, 1 = standard, 2 = rich, 3 = max
	FrontBrakeBias          uint8    // Front brake bias (percentage)
	PitLimiterStatus        uint8    // Pit limiter status - 0 = off, 1 = on
	FuelInTank              float32  // Current fuel mass
	FuelCapacity            float32  // Fuel capacity
	FuelRemainingLaps       float32  // Fuel remaining in terms of laps (value on MFD)
	MaxRPM                  uint16   // Cars max RPM, point of rev limiter
	IdleRPM                 uint16   // Cars idle RPM
	MaxGears                uint8    // Maximum number of gears
	DRSAllowed              uint8    // 0 = not allowed, 1 = allowed, -1 = unknown
	DRSActivationDistance   uint16   // 0 = DRS not available, non-zero - DRS will be available // in [X] metres
	TyresWear               [4]uint8 // Tyre wear percentage
	ActualTyreCompound      uint8    // F1 Modern - 16 = C5, 17 = C4, 18 = C3, 19 = C2, 20 = C1 // 7 = inter, 8 = wet // F1 Classic - 9 = dry, 10 = wet // F2 – 11 = super soft, 12 = soft, 13 = medium, 14 = hard // 15 = wet
	VisualTyreCompound      uint8    // F1 visual (can be different from actual compound) // 16 = soft, 17 = medium, 18 = hard, 7 = inter, 8 = wet // F1 Classic – same as above // F2 – same as above
	TyresAgeLaps            uint8    // Age in laps of the current set of tyres
	TyresDamage             [4]uint8 // Tyre damage (percentage)
	FrontLeftWingDamage     uint8    // Front left wing damage (percentage)
	FrontRightWingDamage    uint8    // Front right wing damage (percentage)
	RearWingDamage          uint8    // Rear wing damage (percentage)
	DRSFault                uint8    // Indicator for DRS fault, 0 = OK, 1 = fault
	EngineDamage            uint8    // Engine damage (percentage)
	GearBoxDamage           uint8    // Gear box damage (percentage)
	VehicleFIAFlags         int8     // -1 = invalid/unknown, 0 = none, 1 = green // 2 = blue, 3 = yellow, 4 = red
	ERSStoreEnergy          float32  // ERS energy store in Joules
	ERSDeployMode           uint8    // ERS deployment mode, 0 = none, 1 = medium // 2 = overtake, 3 = hotlap
	ERSHarvestedThisLapMGUK float32  // ERS energy harvested this lap by MGU-K
	ERSHarvestedThisLapMGUH float32  // ERS energy harvested this lap by MGU-H
	ERSDeployedThisLap      float32  // ERS energy deployed this lap
}

type PacketCarStatusData struct {
	Header        PacketHeader // Header
	CarStatusData [22]CarStatusData
}

type FinalClassificationData struct {
	Position         uint8    // Finishing position
	NumLaps          uint8    // Number of laps completed
	GridPosition     uint8    // Grid position of the car
	Points           uint8    // Number of points scored
	NumPitStops      uint8    // Number of pit stops made
	ResultStatus     uint8    // Result status - 0 = invalid, 1 = inactive, 2 = active // 3 = finished, 4 = disqualified, 5 = not classified // 6 = retired
	BestLapTime      float32  // Best lap time of the session in seconds
	TotalRaceTime    float64  // Total race time in seconds without penalties
	PenaltiesTime    uint8    // Total penalties accumulated in seconds
	NumPenalties     uint8    // Number of penalties applied to this driver
	TyreStints       uint8    // Number of tyres stints up to maximum
	TyreStintsActual [8]uint8 // Actual tyres used by this driver
	TyreStintsVisual [8]uint8 // Visual tyres used by this driver
}

type PacketFinalClassificationData struct {
	Header             PacketHeader // Header
	NumCars            uint8        // Number of cars in the final classification
	ClassificationData [22]FinalClassificationData
}
