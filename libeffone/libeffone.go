package libeffone

type PacketHeaderPacketID uint8

const (
	PacketIDMotion              = 0 // Contains all motion data for player’s car – only sent while player is in control
	PacketIDSession             = 1 // Data about the session – track, time left
	PacketIDLapData             = 2 // Data about all the lap times of cars in the session
	PacketIDEvent               = 3 // Various notable events that happen during a session
	PacketIDParticipants        = 4 // List of participants in the session, mostly relevant for multiplayer
	PacketIDCarSetups           = 5 // Packet detailing car setups for cars in the race
	PacketIDCarTelemetry        = 6 // Telemetry data for all cars
	PacketIDCarStatus           = 7 // Status data for all cars such as damage
	PacketIDFinalClassification = 8 // Final classification confirmation at the end of a race
	PacketIDLobbyInfo           = 9 // Information about players in a multiplayer lobby
)

func (pid PacketHeaderPacketID) String() string {
	switch pid {
	case PacketIDMotion:
		return "Motion"
	case PacketIDSession:
		return "Session"
	case PacketIDLapData:
		return "Lap Data"
	case PacketIDEvent:
		return "Event"
	case PacketIDParticipants:
		return "Participants"
	case PacketIDCarSetups:
		return "Car Setups"
	case PacketIDCarTelemetry:
		return "Car Telemetry"
	case PacketIDCarStatus:
		return "Car Status"
	case PacketIDFinalClassification:
		return "Final Classification"
	case PacketIDLobbyInfo:
		return "Lobby Info"
	}

	return "Unknown packet"
}

type TeamID uint8

const (
	TeamIDMercedes2020            = 0   // Mercedes
	TeamIDFerrari2020             = 1   // Ferrari
	TeamIDRedBull2020             = 2   // Red Bull Racing
	TeamIDWilliams2020            = 3   // Williams
	TeamIDRacingPoint2020         = 4   // Racing Point
	TeamIDRenault2020             = 5   // Renault
	TeamIDAlphaTauri2020          = 6   // Alpha Tauri
	TeamIDHaas2020                = 7   // Haas
	TeamIDMcLaren2020             = 8   // McLaren
	TeamIDAlfaRomeo2020           = 9   // Alfa Romeo
	TeamIDMcLaren1988             = 10  // McLaren 1988
	TeamIDMcLaren1991             = 11  // McLaren 1991
	TeamIDWilliams1992            = 12  // Williams 1992
	TeamIDFerrari1995             = 13  // Ferrari 1995
	TeamIDWilliams1996            = 14  // Williams 1996
	TeamIDMcLaren1998             = 15  // McLaren 1998
	TeamIDFerrari2002             = 16  // Ferrari 2002
	TeamIDFerrari2004             = 17  // Ferrari 2004
	TeamIDRenault2006             = 18  // Renault 2006
	TeamIDFerrari2007             = 19  // Ferrari 2007
	TeamIDMcLaren2008             = 20  // McLaren 2008
	TeamIDRedBull2010             = 21  // Red Bull 2010
	TeamIDFerrari1976             = 22  // Ferrari 1976
	TeamIDARTGrandPrix2020        = 23  // ART Grand Prix
	TeamIDCamposVexatecRacing2020 = 24  // Campos Vexatec Racing
	TeamIDCarlin2020              = 25  // Carlin
	TeamIDCharouzRacingSystem2020 = 26  // Charouz Racing System
	TeamIDDams2020                = 27  // DAMS
	TeamIDRussianTime2020         = 28  // Russian Time
	TeamIDMPMotorsport2020        = 29  // MP Motorsport
	TeamIDPertamina2020           = 30  // Pertamina
	TeamIDMcLaren1990             = 31  // McLaren 1990
	TeamIDTrident2020             = 32  // Trident
	TeamIDBWTArden2020            = 33  // BWT Arden
	TeamIDMcLaren1976             = 34  // McLaren 1976
	TeamIDLotus1972               = 35  // Lotus 1972
	TeamIDFerrari1979             = 36  // Ferrari 1979
	TeamIDMcLaren1982             = 37  // McLaren 1982
	TeamIDWilliams2003            = 38  // Williams 2003
	TeamIDBrawn2009               = 39  // Brawn 2009
	TeamIDLotus1978               = 40  // Lotus 1978
	TeamIDGeneric2020             = 41  // F1 Generic car
	TeamIDArtGP2019               = 42  // Art GP ’19
	TeamIDCampos2019              = 43  // Campos ’19
	TeamIDCarlin2019              = 44  // Carlin ’19
	TeamIDSauberJuniorCharouz2019 = 45  // Sauber Junior Charouz ’19
	TeamIDDams2019                = 46  // Dams ’19
	TeamIDUniVirtuosi2019         = 47  // Uni-Virtuosi ‘19
	TeamIDMPMotorsport2019        = 48  // MP Motorsport ‘19
	TeamIDPrema2019               = 49  // Prema ’19
	TeamIDTrident2019             = 50  // Trident ’19
	TeamIDArden2019               = 51  // Arden ’19
	TeamIDBenetton1994            = 53  // Benetton 1994
	TeamIDBenetton1995            = 54  // Benetton 1995
	TeamIDFerrari2000             = 55  // Ferrari 2000
	TeamIDJordan1991              = 56  // Jordan 1991
	TeamIDMyTeam2020              = 255 // My Team
)

func (tid TeamID) String() string {
	switch tid {
	case TeamIDMercedes2020:
		return "Mercedes"
	case TeamIDFerrari2020:
		return "Ferrari"
	case TeamIDRedBull2020:
		return "Red Bull Racing"
	case TeamIDWilliams2020:
		return "Williams"
	case TeamIDRacingPoint2020:
		return "Racing Point"
	case TeamIDRenault2020:
		return "Renault"
	case TeamIDAlphaTauri2020:
		return "Alpha Tauri"
	case TeamIDHaas2020:
		return "Haas"
	case TeamIDMcLaren2020:
		return "McLaren"
	case TeamIDAlfaRomeo2020:
		return "Alfa Romeo"
	case TeamIDMcLaren1988:
		return "McLaren 1988"
	case TeamIDMcLaren1991:
		return "McLaren 1991"
	case TeamIDWilliams1992:
		return "Williams 1992"
	case TeamIDFerrari1995:
		return "Ferrari 1995"
	case TeamIDWilliams1996:
		return "Williams 1996"
	case TeamIDMcLaren1998:
		return "McLaren 1998"
	case TeamIDFerrari2002:
		return "Ferrari 2002"
	case TeamIDFerrari2004:
		return "Ferrari 2004"
	case TeamIDRenault2006:
		return "Renault 2006"
	case TeamIDFerrari2007:
		return "Ferrari 2007"
	case TeamIDMcLaren2008:
		return "McLaren 2008"
	case TeamIDRedBull2010:
		return "Red Bull 2010"
	case TeamIDFerrari1976:
		return "Ferrari 1976"
	case TeamIDARTGrandPrix2020:
		return "ART Grand Prix"
	case TeamIDCamposVexatecRacing2020:
		return "Campos Vexatec Racing"
	case TeamIDCarlin2020:
		return "Carlin"
	case TeamIDCharouzRacingSystem2020:
		return "Charouz Racing System"
	case TeamIDDams2020:
		return "DAMS"
	case TeamIDRussianTime2020:
		return "Russian Time"
	case TeamIDMPMotorsport2020:
		return "MP Motorsport"
	case TeamIDPertamina2020:
		return "Pertamina"
	case TeamIDMcLaren1990:
		return "McLaren 1990"
	case TeamIDTrident2020:
		return "Trident"
	case TeamIDBWTArden2020:
		return "BWT Arden"
	case TeamIDMcLaren1976:
		return "McLaren 1976"
	case TeamIDLotus1972:
		return "Lotus 1972"
	case TeamIDFerrari1979:
		return "Ferrari 1979"
	case TeamIDMcLaren1982:
		return "McLaren 1982"
	case TeamIDWilliams2003:
		return "Williams 2003"
	case TeamIDBrawn2009:
		return "Brawn 2009"
	case TeamIDLotus1978:
		return "Lotus 1978"
	case TeamIDGeneric2020:
		return "F1 Generic car"
	case TeamIDArtGP2019:
		return "Art GP ’19"
	case TeamIDCampos2019:
		return "Campos ’19"
	case TeamIDCarlin2019:
		return "Carlin ’19"
	case TeamIDSauberJuniorCharouz2019:
		return "Sauber Junior Charouz ’19"
	case TeamIDDams2019:
		return "Dams ’19"
	case TeamIDUniVirtuosi2019:
		return "Uni-Virtuosi ‘19"
	case TeamIDMPMotorsport2019:
		return "MP Motorsport ‘19"
	case TeamIDPrema2019:
		return "Prema ’19"
	case TeamIDTrident2019:
		return "Trident ’19"
	case TeamIDArden2019:
		return "Arden ’19"
	case TeamIDBenetton1994:
		return "Benetton 1994"
	case TeamIDBenetton1995:
		return "Benetton 1995"
	case TeamIDFerrari2000:
		return "Ferrari 2000"
	case TeamIDJordan1991:
		return "Jordan 1991"
	case TeamIDMyTeam2020:
		return "My Team"
	}

	return "Unknown team"
}

type DriverID uint8

const (
	DriverIDCarlosSainz         = 0
	DriverIDDaniilKvyat         = 1
	DriverIDDanielRicciardo     = 2
	DriverIDKimiRaikkonen       = 6
	DriverIDLewisHamilton       = 7
	DriverIDMaxVerstappen       = 9
	DriverIDNicoHulkenberg      = 10
	DriverIDKevinMagnussen      = 11
	DriverIDRomainGrosjean      = 12
	DriverIDSebastianVettel     = 13
	DriverIDSergioPerez         = 14
	DriverIDValtteriBottas      = 15
	DriverIDEstebanOcon         = 17
	DriverIDLanceStroll         = 19
	DriverIDArronBarnes         = 20
	DriverIDMartinGiles         = 21
	DriverIDAlexMurray          = 22
	DriverIDLucasRoth           = 23
	DriverIDIgorCorreia         = 24
	DriverIDSophieLevasseur     = 25
	DriverIDJonasSchiffer       = 26
	DriverIDAlainForest         = 27
	DriverIDJayLetourneau       = 28
	DriverIDEstoSaari           = 29
	DriverIDYasarAtiyeh         = 30
	DriverIDCallistoCalabresi   = 31
	DriverIDNaotaIzum           = 32
	DriverIDHowardClarke        = 33
	DriverIDWilheimKaufmann     = 34
	DriverIDMarieLaursen        = 35
	DriverIDFlavioNieves        = 36
	DriverIDPeterBelousov       = 37
	DriverIDKlimekMichalski     = 38
	DriverIDSantiagoMoreno      = 39
	DriverIDBenjaminCoppens     = 40
	DriverIDNoahVisser          = 41
	DriverIDGertWaldmuller      = 42
	DriverIDJulianQuesada       = 43
	DriverIDDanielJones         = 44
	DriverIDArtemMarkelov       = 45
	DriverIDTadasukeMakino      = 46
	DriverIDSeanGelael          = 47
	DriverIDNyckDeVries         = 48
	DriverIDJackAitken          = 49
	DriverIDGeorgeRussell       = 50
	DriverIDMaximilianGunther   = 51
	DriverIDNireiFukuzumi       = 52
	DriverIDLucaGhiotto         = 53
	DriverIDLandoNorris         = 54
	DriverIDSergioSetteCamara   = 55
	DriverIDLouisDeletraz       = 56
	DriverIDAntonioFuoco        = 57
	DriverIDCharlesLeclerc      = 58
	DriverIDPierreGasly         = 59
	DriverIDAlexanderAlbon      = 62
	DriverIDNicholasLatifi      = 63
	DriverIDDorianBoccolacci    = 64
	DriverIDNikoKari            = 65
	DriverIDRobertoMerhi        = 66
	DriverIDArjunMaini          = 67
	DriverIDAlessioLorandi      = 68
	DriverIDRubenMeijer         = 69
	DriverIDRashidNair          = 70
	DriverIDJackTremblay        = 71
	DriverIDAntonioGiovinazzi   = 74
	DriverIDRobertKubica        = 75
	DriverIDNobuharuMatsushita  = 78
	DriverIDNikitaMazepin       = 79
	DriverIDGuanyaZhou          = 80
	DriverIDMickSchumacher      = 81
	DriverIDCallumIlott         = 82
	DriverIDJualManuelCorrera   = 83
	DriverIDJordanKing          = 84
	DriverIDMahaveerRaghunathan = 85
	DriverIDTatianaCalderon     = 86
	DriverIDAnthoineHubert      = 87
	DriverIDGuilianoAlesi       = 88
	DriverIDRalphBoschung       = 89
)

func (did DriverID) String() string {
	switch did {
	case DriverIDCarlosSainz:
		return "Carlos Sainz"
	case DriverIDDaniilKvyat:
		return "Daniil Kvyat"
	case DriverIDDanielRicciardo:
		return "Daniel Ricciardo"
	case DriverIDKimiRaikkonen:
		return "Kimi Räikkönen"
	case DriverIDLewisHamilton:
		return "Lewis Hamilton"
	case DriverIDMaxVerstappen:
		return "Max Verstappen"
	case DriverIDNicoHulkenberg:
		return "Nico Hulkenburg"
	case DriverIDKevinMagnussen:
		return "Kevin Magnussen"
	case DriverIDRomainGrosjean:
		return "Romain Grosjean"
	case DriverIDSebastianVettel:
		return "Sebastian Vettel"
	case DriverIDSergioPerez:
		return "Sergio Perez"
	case DriverIDValtteriBottas:
		return "Valtteri Bottas"
	case DriverIDEstebanOcon:
		return "Esteban Ocon"
	case DriverIDLanceStroll:
		return "Lance Stroll"
	case DriverIDArronBarnes:
		return "Arron Barnes"
	case DriverIDMartinGiles:
		return "Martin Giles"
	case DriverIDAlexMurray:
		return "Alex Murray"
	case DriverIDLucasRoth:
		return "Lucas Roth"
	case DriverIDIgorCorreia:
		return "Igor Correia"
	case DriverIDSophieLevasseur:
		return "Sophie Levasseur"
	case DriverIDJonasSchiffer:
		return "Jonas Schiffer"
	case DriverIDAlainForest:
		return "Alain Forest"
	case DriverIDJayLetourneau:
		return "Jay Letourneau"
	case DriverIDEstoSaari:
		return "Esto Saari"
	case DriverIDYasarAtiyeh:
		return "Yasar Atiyeh"
	case DriverIDCallistoCalabresi:
		return "Callisto Calabresi"
	case DriverIDNaotaIzum:
		return "Naota Izum"
	case DriverIDHowardClarke:
		return "Howard Clarke"
	case DriverIDWilheimKaufmann:
		return "Wilheim Kaufmann"
	case DriverIDMarieLaursen:
		return "Marie Laursen"
	case DriverIDFlavioNieves:
		return "Flavio Nieves"
	case DriverIDPeterBelousov:
		return "Peter Belousov"
	case DriverIDKlimekMichalski:
		return "Klimek Michalski"
	case DriverIDSantiagoMoreno:
		return "Santiago Moreno"
	case DriverIDBenjaminCoppens:
		return "Benjamin Coppens"
	case DriverIDNoahVisser:
		return "Noah Visser"
	case DriverIDGertWaldmuller:
		return "Gert Waldmuller"
	case DriverIDJulianQuesada:
		return "Julian Quesada"
	case DriverIDDanielJones:
		return "Daniel Jones"
	case DriverIDArtemMarkelov:
		return "Artem Markelov"
	case DriverIDTadasukeMakino:
		return "Tadasuke Makino"
	case DriverIDSeanGelael:
		return "Sean Gelael"
	case DriverIDNyckDeVries:
		return "Nyck De Vries"
	case DriverIDJackAitken:
		return "Jack Aitken"
	case DriverIDGeorgeRussell:
		return "George Russell"
	case DriverIDMaximilianGunther:
		return "Maximilian Günther"
	case DriverIDNireiFukuzumi:
		return "Nirei Fukuzumi"
	case DriverIDLucaGhiotto:
		return "Luca Ghiotto"
	case DriverIDLandoNorris:
		return "Lando Norris"
	case DriverIDSergioSetteCamara:
		return "Sérgio Sette Câmara"
	case DriverIDLouisDeletraz:
		return "Louis Delétraz"
	case DriverIDAntonioFuoco:
		return "Antonio Fuoco"
	case DriverIDCharlesLeclerc:
		return "Charles Leclerc"
	case DriverIDPierreGasly:
		return "Pierre Gasly"
	case DriverIDAlexanderAlbon:
		return "Alexander Albon"
	case DriverIDNicholasLatifi:
		return "Nicholas Latifi"
	case DriverIDDorianBoccolacci:
		return "Dorian Boccolacci"
	case DriverIDNikoKari:
		return "Niko Kari"
	case DriverIDRobertoMerhi:
		return "Roberto Merhi"
	case DriverIDArjunMaini:
		return "Arjun Maini"
	case DriverIDAlessioLorandi:
		return "Alessio Lorandi"
	case DriverIDRubenMeijer:
		return "Ruben Meijer"
	case DriverIDRashidNair:
		return "Rashid Nair"
	case DriverIDJackTremblay:
		return "Jack Tremblay"
	case DriverIDAntonioGiovinazzi:
		return "Antonio Giovinazzi"
	case DriverIDRobertKubica:
		return "Robert Kubica"
	case DriverIDNobuharuMatsushita:
		return "Nobuharu Matsushita"
	case DriverIDNikitaMazepin:
		return "Nikita Mazepin"
	case DriverIDGuanyaZhou:
		return "Guanya Zhou"
	case DriverIDMickSchumacher:
		return "Mick Schumacher"
	case DriverIDCallumIlott:
		return "Callum Ilott"
	case DriverIDJualManuelCorrera:
		return "Juan Manuel Correa"
	case DriverIDJordanKing:
		return "Jordan King"
	case DriverIDMahaveerRaghunathan:
		return "Mahaveer Raghunathan"
	case DriverIDTatianaCalderon:
		return "Tatiana Calderon"
	case DriverIDAnthoineHubert:
		return "Anthoine Hubert 😍"
	case DriverIDGuilianoAlesi:
		return "Guiliano Alesi"
	case DriverIDRalphBoschung:
		return "Ralph Boschung"
	}

	return "Unknown driver"
}

type PacketHeader struct {
	PacketFormat            uint16               // 2020
	GameMajorVersion        uint8                // Game major version - "X.00"
	GameMinorVersion        uint8                // Game minor version - "1.XX"
	PacketVersion           uint8                // Version of this packet type, all start from 1
	PacketID                PacketHeaderPacketID // Identifier for the packet type, see below
	SessionUID              uint64               // Unique identifier for the session
	SessionTime             float32              // Session timestamp
	FrameIdentifier         uint32               // Identifier for the frame the data was retrieved on
	PlayerCarIndex          uint8                // Index of player's car in the array
	SecondaryPlayerCarIndex uint8                // Index of secondary player's car in the array (splitscreen) // 255 if no second player
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

type ZoneFlag int8

const (
	ZoneFlagUnknown = -1
	ZoneFlagNone    = 0
	ZoneFlagGreen   = 1
	ZoneFlagBlue    = 2
	ZoneFlagYellow  = 3
	ZoneFlagRed     = 4
)

func (zf ZoneFlag) String() string {
	switch zf {
	case ZoneFlagUnknown:
		return "Unknown"
	case ZoneFlagNone:
		return "None"
	case ZoneFlagGreen:
		return "Green"
	case ZoneFlagBlue:
		return "Blue"
	case ZoneFlagYellow:
		return "Yellow"
	case ZoneFlagRed:
		return "Red"
	}

	return "Unknown"
}

type MarshalZone struct {
	ZoneStart float32  // Fraction (0..1) of way through the lap the marshal zone starts
	ZoneFlag  ZoneFlag // -1 = invalid/unknown, 0 = none, 1 = green, 2 = blue, 3 = yellow, 4 = red
}

type SessionType uint8

const (
	SessionTypeUnknown = 0
	SessionTypeP1      = 1
	SessionTypeP2      = 2
	SessionTypeP3      = 3
	SessionTypeShortP  = 4
	SessionTypeQ1      = 5
	SessionTypeQ2      = 6
	SessionTypeQ3      = 7
	SessionTypeShortQ  = 8
	SessionTypeOSQ     = 9
	SessionTypeR       = 10
	SessionTypeR2      = 11
	SessionTypeTT      = 12
)

func (st SessionType) String() string {
	switch st {
	case SessionTypeUnknown:
		return "Unknown"
	case SessionTypeP1:
		return "Practice 1"
	case SessionTypeP2:
		return "Practice 2"
	case SessionTypeP3:
		return "Practice 3"
	case SessionTypeShortP:
		return "Short Practice"
	case SessionTypeQ1:
		return "Qualifier 1"
	case SessionTypeQ2:
		return "Qualifier 2"
	case SessionTypeQ3:
		return "Qualifier 3"
	case SessionTypeShortQ:
		return "Short Qualifier"
	case SessionTypeOSQ:
		return "One-shot Qualifier"
	case SessionTypeR:
		return "Race"
	case SessionTypeR2:
		return "Race 2"
	case SessionTypeTT:
		return "Time Trial"
	}

	return "Unknown"
}

type WeatherForecastSample struct {
	SessionType      SessionType // 0 = unknown, 1 = P1, 2 = P2, 3 = P3, 4 = Short P, 5 = Q1 // 6 = Q2, 7 = Q3, 8 = Short Q, 9 = OSQ, 10 = R, 11 = R2 // 12 = Time Trial
	TimeOffset       uint8       // Time in minutes the forecast is for
	Weather          uint8       // Weather - 0 = clear, 1 = light cloud, 2 = overcast // 3 = light rain, 4 = heavy rain, 5 = storm
	TrackTemperature int8        // Track temp. in degrees celsius
	AirTemperature   int8        // Air temp. in degrees celsius
}

type PacketSessionData struct {
	Header                    PacketHeader              // Header
	Weather                   uint8                     // Weather - 0 = clear, 1 = light cloud, 2 = overcast // 3 = light rain, 4 = heavy rain, 5 = storm
	TrackTemperature          int8                      // Track temp. in degrees celsius
	AirTemperature            int8                      // Air temp. in degrees celsius
	TotalLaps                 uint8                     // Total number of laps in this race
	TrackLength               uint16                    // Track length in metres
	SessionType               SessionType               // 0 = unknown, 1 = P1, 2 = P2, 3 = P3, 4 = Short P // 5 = Q1, 6 = Q2, 7 = Q3, 8 = Short Q, 9 = OSQ // 10 = R, 11 = R2, 12 = Time Trial
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
	DriverID      DriverID  // Driver id - see appendix
	TeamID        TeamID    // Team id - see appendix
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

type LobbyInfoData struct {
	AIControlled uint8     // Whether the vehicle is AI (1) or Human (0) controlled
	TeamID       TeamID    // Team id - see appendix (255 if no team currently selected)
	Nationality  uint8     // Nationality of the driver
	Name         [48]uint8 // Name of participant in UTF-8 format – null terminated // Will be truncated with ... (U+2026) if too long
	ReadyStatus  uint8     // 0 = not ready, 1 = ready, 2 = spectating
}

type PacketLobbyInfoData struct {
	Header       PacketHeader // Header
	NumPlayers   uint8        // Number of players in the lobby data
	LobbyPlayers [22]LobbyInfoData
}
