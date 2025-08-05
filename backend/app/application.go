package app

import (
	"fmt"
	"log"

	domainrepo "github.com/joaopedropio/musiquera/app/domain/repo"
	infra "github.com/joaopedropio/musiquera/app/infra"
	"github.com/joaopedropio/musiquera/app/utils"

	"github.com/jmoiron/sqlx"
	//_ "github.com/mattn/go-sqlite3"
	_ "modernc.org/sqlite"

)

type Application interface {
	DBConnection() *sqlx.DB
	Close() error
	LoginService() infra.LoginService
	PasswordService() infra.PasswordService
	Repo() domainrepo.Repo
	Environment() Environment
	UserRepo() infra.UserRepo
}

type application struct {
	db              *sqlx.DB
	repo            domainrepo.Repo
	env             Environment
	passwordService infra.PasswordService
	userRepo        infra.UserRepo
	loginService    infra.LoginService
}

func (a *application) Environment() Environment {
	return a.env
}

func (a *application) Repo() domainrepo.Repo {
	return a.repo
}

func (a *application) PasswordService() infra.PasswordService {
	return a.passwordService
}

func (a *application) UserRepo() infra.UserRepo {
	return a.userRepo
}

func (a *application) LoginService() infra.LoginService {
	return a.loginService
}

func (a *application) DBConnection() *sqlx.DB {
	return a.db
}

func (a *application) Close() error {
	fmt.Println("closing db connection")
	if err := a.db.Close(); err != nil {
		log.Fatalf("unable to close db connection: %s", err)
	}
	return nil
}

func NewApplication() (Application, error) {
	env := GetEnvironmentVariables()
	db, err := sqlx.Open("sqlite", env.DatabaseDir+"/musiquera.db?_foreign_keys=on")
	if err != nil {
		panic(fmt.Errorf("unable to start db connection: %w", err))
	}
	repo := infra.NewRepo(db)
	passwordService := infra.NewPasswordService(env.JWTSecret)
	userRepo := infra.NewUserRepo(db)
	loginService := infra.NewLoginService(passwordService, userRepo)
	a := &application{
		db,
		repo,
		env,
		passwordService,
		userRepo,
		loginService,
	}
	a.schema(db)
	/*
	if err := a.feed(); err != nil {
		return nil, fmt.Errorf("unable to feed: %w", err)
	}
	*/
	return a, nil
}

func (a *application) schema(db *sqlx.DB) {
	db.MustExec(utils.DatabaseSchema())
}

/*
func (a *application) feed() error {

	zef := domain.NewArtist(uuid.New(), "Danimal Cannon", "/media/danimal_cannon_and_zef/danimal_cannon_profile.jpg", time.Now())
	err := a.repo.AddArtist(zef)
	if err != nil {
		return fmt.Errorf("unable to add artist: %w", err)
	}
	parallelProcessing := []domain.Track{
		domain.NewTrack(uuid.New(), "Logic Gatekeeper", "", "/media/danimal_cannon_and_zef/parallel_processing/Danimal_Cannon___Zef___Logic_Gatekeeper__uLtRVm_tkxI_/manifest.mpd", time.Minute*5, nil, time.Now()),
		domain.NewTrack(uuid.New(), "Legacy", "", "/media/danimal_cannon_and_zef/parallel_processing/Danimal_Cannon___Zef___Legacy__pD3JoYx9hfw_/manifest.mpd", time.Minute*5, nil, time.Now()),
		domain.NewTrack(uuid.New(), "Chronos", "", "/media/danimal_cannon_and_zef/parallel_processing/Danimal_Cannon___Zef___Chronos__eLG_K7X2BaE_/manifest.mpd", time.Minute*5, nil, time.Now()),
		domain.NewTrack(uuid.New(), "Rhapsody", "", "/media/danimal_cannon_and_zef/parallel_processing/Danimal_Cannon___Zef___Rhapsody__5MSeUK93_Sk_/manifest.mpd", time.Minute*5, nil, time.Now()),
		domain.NewTrack(uuid.New(), "Glitch", "", "/media/danimal_cannon_and_zef/parallel_processing/Danimal_Cannon___Zef___Glitch__upX7bibwJBw_/manifest.mpd", time.Minute*5, nil, time.Now()),
		domain.NewTrack(uuid.New(), "The Lunar Whale", "", "/media/danimal_cannon_and_zef/parallel_processing/Danimal_Cannon___Zef___The_Lunar_Whale__p1XpTZPILXM_/manifest.mpd", time.Minute*5, nil, time.Now()),
		domain.NewTrack(uuid.New(), "Corrupted", "", "/media/danimal_cannon_and_zef/parallel_processing/Danimal_Cannon___Zef___Corrupted__6j_eb_8huQc_/manifest.mpd", time.Minute*5, nil, time.Now()),
		domain.NewTrack(uuid.New(), "Parallel Processing", "", "/media/danimal_cannon_and_zef/parallel_processing/Danimal_Cannon___Zef___Parallel_Processing__1Fuch5gfUq0_/manifest.mpd", time.Minute*5, nil, time.Now()),
	}

	err = a.repo.AddFullRelease(domain.NewFullRelease(
		uuid.New(),
		"Parallel Processing",
		domain.ReleaseTypeAlbum,
		"/media/danimal_cannon_and_zef/parallel_processing/parallel_processing_cover.jpg",
		domain.NewDate(2013, 1, 15),
		zef,
		parallelProcessing,
		time.Now(),
	))
	if err != nil {
		return fmt.Errorf("unable to add release: %w", err)
	}

	kream := domain.NewArtist(uuid.New(), "Kream", "/media/kream/kream_profile_cover.jpg", time.Now())
	err = a.repo.AddArtist(kream)
	if err != nil {
		return fmt.Errorf("unable to add artist: %w", err)
	}
	
	liquidLabVol12TrackID := uuid.New()
	liquidLabVol12 := []domain.Track{
		domain.NewTrack(liquidLabVol12TrackID, "Live set", "", "/media/kream/liquid_lab_vol_12/KREAM_Presents___LIQUID___LAB_Vol__12__Adriatique__Prospa__WhoMadeWho__at_Vemork__LUeHFSkbLlo_/manifest.mpd", time.Minute*5, []domain.Segment{
			domain.NewSegment(liquidLabVol12TrackID, "KREAM - Liquid Lab Intro", 67),
			domain.NewSegment(liquidLabVol12TrackID, "Dyzen - She Likes vs", 82),
			domain.NewSegment(liquidLabVol12TrackID, "RÜFÜS DU SOL - Inhale (SCRIPT Edit) vs", 128),
			domain.NewSegment(liquidLabVol12TrackID, "RÜFÜS DU SOL - Lately", 189),
			domain.NewSegment(liquidLabVol12TrackID, "NORRA - This Love", 265),
			domain.NewSegment(liquidLabVol12TrackID, "RY X - Only vs", 325),
			domain.NewSegment(liquidLabVol12TrackID, "Simon Doty & DJ Pierre - Come Together", 372),
			domain.NewSegment(liquidLabVol12TrackID, "Dom Dolla ft Daya - Dreamin’ (KREAM Remix)", 449),
			domain.NewSegment(liquidLabVol12TrackID, "KREAM & Volaris - ID (Set You Free)", 713),
			domain.NewSegment(liquidLabVol12TrackID, "Vintage Culture & Tom Breu ft Maverick Sabre - Weak (Acappella) vs", 894),
			domain.NewSegment(liquidLabVol12TrackID, "Cristoph & Harry Diamond - Hold Me Close vs", 946),
			domain.NewSegment(liquidLabVol12TrackID, "Florence + The Machine - You've Got The Love (Acappella)", 1022),
			domain.NewSegment(liquidLabVol12TrackID, "Aaron Hibell ft Felsmann + Tiley - Levitation vs", 1235),
			domain.NewSegment(liquidLabVol12TrackID, "Elderbrook ft George FitzGerald - Glad I Found You (Acappella) vs", 1280),
			domain.NewSegment(liquidLabVol12TrackID, "Morgin Madison & Adam Nazar vs Crimsen - Closer (Ryan Lucian Remix)", 1341),
			domain.NewSegment(liquidLabVol12TrackID, "Prospa ft RAHH - This Rhythm (Acappella) vs", 1418),
			domain.NewSegment(liquidLabVol12TrackID, "Syence - The Distance", 1448),
			domain.NewSegment(liquidLabVol12TrackID, "Eric Prydz - 2night vs", 1616),
			domain.NewSegment(liquidLabVol12TrackID, "Emmit Fenn - The Chase (Acappella)", 1632),
			domain.NewSegment(liquidLabVol12TrackID, "Adriatique & WhoMadeWho - Miracle (RÜFÜS DU SOL Remix)", 1920),
			domain.NewSegment(liquidLabVol12TrackID, "KREAM & Ruback - Se Que Quiere", 2073),
			domain.NewSegment(liquidLabVol12TrackID, "Silar - The Tunnel vs Pa Salieu ft Obongjayar - Style & Fashion (Acappella) vs", 2349),
			domain.NewSegment(liquidLabVol12TrackID, "Dan Sushi - Orbital", 2418),
			domain.NewSegment(liquidLabVol12TrackID, "Son Of Son - Lost Control (Acappella) vs", 2494),
			domain.NewSegment(liquidLabVol12TrackID, "Samm & Ajna - Move vs", 2525),
			domain.NewSegment(liquidLabVol12TrackID, "PACS - Hyperdrive", 2615),
			domain.NewSegment(liquidLabVol12TrackID, "Lil Yachty ft Future & Playboi Carti - Flex Up (Acappella) vs", 2693),
			domain.NewSegment(liquidLabVol12TrackID, "Goom Gum - Staccato", 2738),
			domain.NewSegment(liquidLabVol12TrackID, "Fideles ft Be No Rain - Night After Night (CamelPhat Remix) vs", 2942),
			domain.NewSegment(liquidLabVol12TrackID, "Anakim & Dark Heart vs Frýnn - Seconds Away vs", 2998),
			domain.NewSegment(liquidLabVol12TrackID, "Diplo & HUGEL ft Malou vs Yuna - Forever (Acappella) vs", 3039),
			domain.NewSegment(liquidLabVol12TrackID, "Adriatique & Solique vs ALSO ASTIR - Changing Colors", 3099),
			domain.NewSegment(liquidLabVol12TrackID, "Dyzen - Try vs", 3177),
			domain.NewSegment(liquidLabVol12TrackID, "Cristoph & Pete Tong ft Paul Rogers - Where's The Music Gone", 3221),
			domain.NewSegment(liquidLabVol12TrackID, "KREAM - Manta", 3390),
			domain.NewSegment(liquidLabVol12TrackID, "Rivo - Last Night vs", 3696),
			domain.NewSegment(liquidLabVol12TrackID, "Nils Hoffmann - Closer (Simon Doty Remix)", 3740),
		}, time.Now()),
	}

	err = a.repo.AddFullRelease(domain.NewFullRelease(
		uuid.New(),
		"Liquid Lab Vol 12",
		domain.ReleaseTypeLiveSet,
		"/media/kream/liquid_lab_vol_12/liquid_lab_vol_12_cover.jpg",
		domain.NewDate(2025, 7, 5),
		kream,
		liquidLabVol12,
		time.Now(),
	))
	if err != nil {
		return fmt.Errorf("unable to add release: %w", err)
	}

	massacration := domain.NewArtist(uuid.New(), "Massacration", "/media/massacration/massacration_profile_cover.webp", time.Now())
	err = a.repo.AddArtist(massacration)
	if err != nil {
		return fmt.Errorf("unable to add artist: %w", err)
	}

	gatesOfMetalFriedChickenOfDeath := []domain.Track{
		domain.NewTrack(uuid.New(), "Away Doom", "", "/media/massacration/gates_of_metal_fried_chicken_of_death/Massacration___Away_Doom__QZurWCQGSkA_/manifest.mpd", time.Minute*5, nil, time.Now()),
		domain.NewTrack(uuid.New(), "Cereal Metal", "", "/media/massacration/gates_of_metal_fried_chicken_of_death/Massacration___Cereal_Metal__JFP9dUUGtAo_/manifest.mpd", time.Minute*5, nil, time.Now()),
		domain.NewTrack(uuid.New(), "Evil Papagali", "", "/media/massacration/gates_of_metal_fried_chicken_of_death/Massacration___Evil_Papagali__VTVYAqKdur0_/manifest.mpd", time.Minute*5, nil, time.Now()),
		domain.NewTrack(uuid.New(), "Feel The Fire From Barbecue", "", "/media/massacration/gates_of_metal_fried_chicken_of_death/Massacration___Feel_The_Fire_From_Barbecue__xH7VsEcZtN4_/manifest.mpd", time.Minute*5, nil, time.Now()),
		domain.NewTrack(uuid.New(), "Intro", "", "/media/massacration/gates_of_metal_fried_chicken_of_death/Massacration___Intro__dLylOkbplPQ_/manifest.mpd", time.Minute*5, nil, time.Now()),
		domain.NewTrack(uuid.New(), "Let's Ride To Metal Land The Passage is R$1,0", "", "/media/massacration/gates_of_metal_fried_chicken_of_death/Massacration___Let_s_Ride_To_Metal_Land_The_Passage_is_R_1_0__wGhigjK8H2c_/manifest.mpd", time.Minute*5, nil, time.Now()),
		domain.NewTrack(uuid.New(), "Metal Bucetation", "", "/media/massacration/gates_of_metal_fried_chicken_of_death/Massacration___Metal_Bucetation__rlft_Ff4cfQ_/manifest.mpd", time.Minute*5, nil, time.Now()),
		domain.NewTrack(uuid.New(), "Metal Dental Destruction", "", "/media/massacration/gates_of_metal_fried_chicken_of_death/Massacration___Metal_Dental_Destruction__FjIQlPzTcGs_/manifest.mpd", time.Minute*5, nil, time.Now()),
		domain.NewTrack(uuid.New(), "Metal Glu Glu", "", "/media/massacration/gates_of_metal_fried_chicken_of_death/Massacration___Metal_Glu_Glu__WUtdTBeg11Y_/manifest.mpd", time.Minute*5, nil, time.Now()),
		domain.NewTrack(uuid.New(), "Metal Is The Law", "", "/media/massacration/gates_of_metal_fried_chicken_of_death/Massacration___Metal_Is_The_Law__VVKxhOosHBg_/manifest.mpd", time.Minute*5, nil, time.Now()),
		domain.NewTrack(uuid.New(), "Metal Massacre Attack Aruê Aruô", "", "/media/massacration/gates_of_metal_fried_chicken_of_death/Massacration___Metal_Massacre_Attack_Aru__Aru___g6lQl5JjbmI_/manifest.mpd", time.Minute*5, nil, time.Now()),
		domain.NewTrack(uuid.New(), "Metal Milkshake", "", "/media/massacration/gates_of_metal_fried_chicken_of_death/Massacration___Metal_Milkshake__TMgMVrnOYEw_/manifest.mpd", time.Minute*5, nil, time.Now()),
		domain.NewTrack(uuid.New(), "The God Master", "", "/media/massacration/gates_of_metal_fried_chicken_of_death/Massacration___The_God_Master__3WQqnb0jqpM_/manifest.mpd", time.Minute*5, nil, time.Now()),
	}

	err = a.repo.AddFullRelease(domain.NewFullRelease(
		uuid.New(),
		"Gates of Metal Fried Chicken of Death",
		domain.ReleaseTypeAlbum,
		"/media/massacration/gates_of_metal_fried_chicken_of_death/gates_of_metal_fried_chicken_of_death.jpg",
		domain.NewDate(2008, 3, 7),
		massacration,
		gatesOfMetalFriedChickenOfDeath,
		time.Now(),
	))
	if err != nil {
		return fmt.Errorf("unable to add release: %w", err)
	}

	meshuggah := domain.NewArtist(uuid.New(), "Meshuggah", "/media/meshuggah/meshuggah_profile_cover.jpg", time.Now())
	err = a.repo.AddArtist(meshuggah)
	if err != nil {
		return fmt.Errorf("unable to add artist: %w", err)
	}

	theOphidianTrek := []domain.Track{
		domain.NewTrack(uuid.New(), "Bleed", "", "/media/meshuggah/the_ophidian_trek/Bleed__Live___0F_zJk2oa_w_/manifest.mpd", time.Minute*5, nil, time.Now()),
		domain.NewTrack(uuid.New(), "Combustion", "", "/media/meshuggah/the_ophidian_trek/Combustion__Live___LaMUcYtV_3M_/manifest.mpd", time.Minute*5, nil, time.Now()),
		domain.NewTrack(uuid.New(), "Dancers to a Discordant System", "", "/media/meshuggah/the_ophidian_trek/Dancers_to_a_Discordant_System__Live___AnPJKuRasEU_/manifest.mpd", time.Minute*5, nil, time.Now()),
		domain.NewTrack(uuid.New(), "Demiurge", "", "/media/meshuggah/the_ophidian_trek/Demiurge__Live___ET6BOPI9mpY_/manifest.mpd", time.Minute*5, nil, time.Now()),
		domain.NewTrack(uuid.New(), "Do Not Look Down", "", "/media/meshuggah/the_ophidian_trek/Do_Not_Look_Down__Live___PCTYWKqUULE_/manifest.mpd", time.Minute*5, nil, time.Now()),
		domain.NewTrack(uuid.New(), "I Am Colossus", "", "/media/meshuggah/the_ophidian_trek/I_Am_Colossus__Live___FZrTZLaoP_Y_/manifest.mpd", time.Minute*5, nil, time.Now()),
		domain.NewTrack(uuid.New(), "Lethargica", "", "/media/meshuggah/the_ophidian_trek/Lethargica__Live___CwxAtsPXozc_/manifest.mpd", time.Minute*5, nil, time.Now()),
		domain.NewTrack(uuid.New(), "Mirrors / In Death – Is Life / In Death – Is Death", "", "/media/meshuggah/the_ophidian_trek/Mind_s_Mirrors___In_Death___Is_Life___In_Death___Is_Death__Live___FNn0z2lPZqw_/manifest.mpd", time.Minute*5, nil, time.Now()),
		domain.NewTrack(uuid.New(), "New_Millennium_Cyanide_Christ", "", "/media/meshuggah/the_ophidian_trek/New_Millennium_Cyanide_Christ__Live___EKWQvz0OYtY_/manifest.mpd", time.Minute*5, nil, time.Now()),
		domain.NewTrack(uuid.New(), "Obzen", "", "/media/meshuggah/the_ophidian_trek/Obzen__Live___spZ2GT3kM1U_/manifest.mpd", time.Minute*5, nil, time.Now()),
		domain.NewTrack(uuid.New(), "Rational Gaze", "", "/media/meshuggah/the_ophidian_trek/Rational_Gaze__Live___rV_GzVzdWxk_/manifest.mpd", time.Minute*5, nil, time.Now()),
		domain.NewTrack(uuid.New(), "Swarm", "", "/media/meshuggah/the_ophidian_trek/Swarm__Live___FaOdBys4pgQ_/manifest.mpd", time.Minute*5, nil, time.Now()),
		domain.NewTrack(uuid.New(), "Swarmer", "", "/media/meshuggah/the_ophidian_trek/Swarmer__Live___1HhNYMD1_Zk_/manifest.mpd", time.Minute*5, nil, time.Now()),
		domain.NewTrack(uuid.New(), "The Hurt That Finds You First", "", "/media/meshuggah/the_ophidian_trek/The_Hurt_That_Finds_You_First__Live___xm7jBTRk4p4_/manifest.mpd", time.Minute*5, nil, time.Now()),
		domain.NewTrack(uuid.New(), "The Last Vigil", "", "/media/meshuggah/the_ophidian_trek/The_Last_Vigil__Live___FDWRJQuINnk_/manifest.mpd", time.Minute*5, nil, time.Now()),
	}

	err = a.repo.AddFullRelease(domain.NewFullRelease(
		uuid.New(),
		"The Ophidian Trek (Live)",
		domain.ReleaseTypeAlbum,
		"/media/meshuggah/the_ophidian_trek/ophidian_trek_cover.jpg",
		domain.NewDate(2014, 9, 29),
		meshuggah,
		theOphidianTrek,
		time.Now(),
	))
	if err != nil {
		return fmt.Errorf("unable to add release: %w", err)
	}

	obzen := []domain.Track{
		domain.NewTrack(uuid.New(), "Bleed", "", "/media/meshuggah/obzen/Bleed__GAulPs96ass_/manifest.mpd", time.Minute*5, nil, time.Now()),
		domain.NewTrack(uuid.New(), "Combustion", "", "/media/meshuggah/obzen/Combustion__RL7RFrInBww_/manifest.mpd", time.Minute*5, nil, time.Now()),
		domain.NewTrack(uuid.New(), "Dancers_to_a_Discordant_System", "", "/media/meshuggah/obzen/Dancers_to_a_Discordant_System__0DxI_ZOGbXo_/manifest.mpd", time.Minute*5, nil, time.Now()),
		domain.NewTrack(uuid.New(), "Electric_Red", "", "/media/meshuggah/obzen/Electric_Red__UpTmMSXm9rw_/manifest.mpd", time.Minute*5, nil, time.Now()),
		domain.NewTrack(uuid.New(), "Lethargica", "", "/media/meshuggah/obzen/Lethargica__Qywg2ZjdqMo_/manifest.mpd", time.Minute*5, nil, time.Now()),
		domain.NewTrack(uuid.New(), "Obzen", "", "/media/meshuggah/obzen/Obzen__Fc7aeQGLccI_/manifest.mpd", time.Minute*5, nil, time.Now()),
		domain.NewTrack(uuid.New(), "Pineal_Gland_Optics", "", "/media/meshuggah/obzen/Pineal_Gland_Optics__VgD2Ks_gxxw_/manifest.mpd", time.Minute*5, nil, time.Now()),
		domain.NewTrack(uuid.New(), "Pravus", "", "/media/meshuggah/obzen/Pravus__6KJ2RCm_bcs_/manifest.mpd", time.Minute*5, nil, time.Now()),
		domain.NewTrack(uuid.New(), "This_Spiteful_Snake", "", "/media/meshuggah/obzen/This_Spiteful_Snake__cjjOW5X8nIw_/manifest.mpd", time.Minute*5, nil, time.Now()),
	}
	err = a.repo.AddFullRelease(domain.NewFullRelease(
		uuid.New(),
		"Obzen",
		domain.ReleaseTypeAlbum,
		"/media/meshuggah/obzen/obsen_cover.jpg",
		domain.NewDate(2008, 3, 7),
		meshuggah,
		obzen,
		time.Now(),
	))
	if err != nil {
		return fmt.Errorf("unable to add release: %w", err)
	}

	nothingRemastered := []domain.Track{
		domain.NewTrack(uuid.New(), "Closed Eye Visuals", "", "/media/meshuggah/nothing_remastered/Closed_Eye_Visuals__Remastered_2006___kZ84SjwKe80_/manifest.mpd", time.Minute*5, nil, time.Now()),
		domain.NewTrack(uuid.New(), "Glints Collide", "", "/media/meshuggah/nothing_remastered/Glints_Collide__Remastered_2006___rU6oJ0Kdews_/manifest.mpd", time.Minute*5, nil, time.Now()),
		domain.NewTrack(uuid.New(), "Nebulous", "", "/media/meshuggah/nothing_remastered/Nebulous__Remastered_2006___K6ncYs9Bji8_/manifest.mpd", time.Minute*5, nil, time.Now()),
		domain.NewTrack(uuid.New(), "Obsidian", "", "/media/meshuggah/nothing_remastered/Obsidian__Remastered_2006___TOkPHeGuWuQ_/manifest.mpd", time.Minute*5, nil, time.Now()),
		domain.NewTrack(uuid.New(), "Organic Shadows", "", "/media/meshuggah/nothing_remastered/Organic_Shadows__Remastered_2006___JDlJB_ClTdg_/manifest.mpd", time.Minute*5, nil, time.Now()),
		domain.NewTrack(uuid.New(), "Perpetual Black Second", "", "/media/meshuggah/nothing_remastered/Perpetual_Black_Second__Remastered_2006____NkPYAU1I1U_/manifest.mpd", time.Minute*5, nil, time.Now()),
		domain.NewTrack(uuid.New(), "Rational Gaze", "", "/media/meshuggah/nothing_remastered/Rational_Gaze__Remastered_2006___IcdlMBvMCjs_/manifest.mpd", time.Minute*5, nil, time.Now()),
		domain.NewTrack(uuid.New(), "Spasm", "", "/media/meshuggah/nothing_remastered/Spasm__Remastered_2006___aZAYKpYkHNI_/manifest.mpd", time.Minute*5, nil, time.Now()),
		domain.NewTrack(uuid.New(), "Stengah", "", "/media/meshuggah/nothing_remastered/Stengah__Remastered_2006___ntISKKjf0gk_/manifest.mpd", time.Minute*5, nil, time.Now()),
		domain.NewTrack(uuid.New(), "Straws Pulled at Random", "", "/media/meshuggah/nothing_remastered/Straws_Pulled_at_Random__Remastered_2006___xA_qx8ht2j0_/manifest.mpd", time.Minute*5, nil, time.Now()),
	}
	err = a.repo.AddFullRelease(domain.NewFullRelease(
		uuid.New(),
		"Nothing (Remastered 2006)",
		domain.ReleaseTypeAlbum,
		"/media/meshuggah/nothing_remastered/nothing_cover.jpg",
		domain.NewDate(2006, 10, 31),
		meshuggah,
		nothingRemastered,
		time.Now(),
	))
	if err != nil {
		return fmt.Errorf("unable to add release: %w", err)
	}

	metallica := domain.NewArtist(uuid.New(), "Metallica", "/media/metallica/metallica_profile_cover.jpg", time.Now())
	err = a.repo.AddArtist(metallica)
	if err != nil {
		return fmt.Errorf("unable to add artist: %w", err)
	}
	ridingTheLigthningSongs := []domain.Track{
		domain.NewTrack(uuid.New(), "Creeping Death", "", "/media/metallica/riding_the_lightning/Creeping_Death__Studio_Version___2h4iqDSzVv0_/manifest.mpd", time.Minute*5, nil, time.Now()),
		domain.NewTrack(uuid.New(), "Escape", "", "/media/metallica/riding_the_lightning/Escape__Studio_Version___MGdKPy98Byg_/manifest.mpd", time.Minute*5, nil, time.Now()),
		domain.NewTrack(uuid.New(), "Fade To Black", "", "/media/metallica/riding_the_lightning/Fade_To_Black__Studio_Version___eDZPLSexHVM_/manifest.mpd", time.Minute*5, nil, time.Now()),
		domain.NewTrack(uuid.New(), "Fight Fire With Fire", "", "/media/metallica/riding_the_lightning/Fight_Fire_With_Fire__Studio_Version___ZnCFWlso_UQ_/manifest.mpd", time.Minute*5, nil, time.Now()),
		domain.NewTrack(uuid.New(), "For Whom The Bell Tolls", "", "/media/metallica/riding_the_lightning/For_Whom_The_Bell_Tolls__Studio_Version____b6tJMD34qw_/manifest.mpd", time.Minute*5, nil, time.Now()),
		domain.NewTrack(uuid.New(), "Ride The Lightning", "", "/media/metallica/riding_the_lightning/Ride_The_Lightning__Studio_Version___ArgdUZKslPw_/manifest.mpd", time.Minute*5, nil, time.Now()),
		domain.NewTrack(uuid.New(), "The Call Of Ktulu", "", "/media/metallica/riding_the_lightning/The_Call_Of_Ktulu__Studio_Version___Z_wccw663BE_/manifest.mpd", time.Minute*5, nil, time.Now()),
		domain.NewTrack(uuid.New(), "Trapped Under Ice", "", "/media/metallica/riding_the_lightning/Trapped_Under_Ice__Studio_Version___6mLDoLWJKZw_/manifest.mpd", time.Minute*5, nil, time.Now()),
	}

	err = a.repo.AddFullRelease(domain.NewFullRelease(
		uuid.New(),
		"Riding the Lightning",
		domain.ReleaseTypeAlbum,
		"/media/metallica/riding_the_lightning/riding_the_lightning_cover.jpg",
		domain.NewDate(1984, 7, 27),
		metallica,
		ridingTheLigthningSongs,
		time.Now(),
	))
	if err != nil {
		return fmt.Errorf("unable to add release: %w", err)
	}

	masterOfPuppetsSongs := []domain.Track{
		domain.NewTrack(uuid.New(), "Battery", "", "/media/metallica/master_of_puppets/Battery__Remastered___uzlOcupu5UE_/manifest.mpd", time.Minute*5, nil, time.Now()),
		domain.NewTrack(uuid.New(), "Damage Inc.", "", "/media/metallica/master_of_puppets/Damage__Inc___Remastered___Abe3AZhcGQs_/manifest.mpd", time.Minute*5, nil, time.Now()),
		domain.NewTrack(uuid.New(), "Disposable Heroes", "", "/media/metallica/master_of_puppets/Disposable_Heroes__Remastered___p3Y8VSVyYN8_/manifest.mpd", time.Minute*5, nil, time.Now()),
		domain.NewTrack(uuid.New(), "Leper Messiah", "", "/media/metallica/master_of_puppets/Leper_Messiah__Remastered___dJp5r4HdRn4_/manifest.mpd", time.Minute*5, nil, time.Now()),
		domain.NewTrack(uuid.New(), "Master Of Puppets", "", "/media/metallica/master_of_puppets/Master_Of_Puppets__Remastered___u6LahTuw02c_/manifest.mpd", time.Minute*5, nil, time.Now()),
		domain.NewTrack(uuid.New(), "Orion", "", "/media/metallica/master_of_puppets/Orion__Remastered___z7bUJPj_8v0_/manifest.mpd", time.Minute*5, nil, time.Now()),
		domain.NewTrack(uuid.New(), "The Thing That Should Never Be", "", "/media/metallica/master_of_puppets/The_Thing_That_Should_Not_Be__Remastered___gm9c_QpuMms_/manifest.mpd", time.Minute*5, nil, time.Now()),
		domain.NewTrack(uuid.New(), "Welcome Home (Sanitarium)", "", "/media/metallica/master_of_puppets/Welcome_Home__Sanitarium___Remastered___G_868UwoJvM_/manifest.mpd", time.Minute*5, nil, time.Now()),
	}

	err = a.repo.AddFullRelease(domain.NewFullRelease(
		uuid.New(),
		"Master of Puppets",
		domain.ReleaseTypeAlbum,
		"/media/metallica/master_of_puppets/master_of_puppets_cover.jpg",
		domain.NewDate(1986, 3, 3),
		metallica,
		masterOfPuppetsSongs,
		time.Now(),
	))
	if err != nil {
		return fmt.Errorf("unable to add release: %w", err)
	}
	return nil
}
*/
