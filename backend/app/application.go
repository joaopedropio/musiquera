package app

import (
	"fmt"
	"time"

	domain "github.com/joaopedropio/musiquera/app/domain/entity"
	domainrepo "github.com/joaopedropio/musiquera/app/domain/repo"
	infra "github.com/joaopedropio/musiquera/app/infra"
)

type Application interface {
	Repo() domainrepo.Repo
	Environment() Environment
}

type application struct {
	repo domainrepo.Repo
	env  Environment
}

func (a *application) Environment() Environment {
	return a.env
}

func (a *application) Repo() domainrepo.Repo {
	return a.repo
}

func NewApplication() (Application, error) {
	repo := infra.NewRepo()
	env := GetEnvironmentVariables()
	a := &application{
		repo,
		env,
	}
	if err := a.feed(); err != nil {
		return nil, fmt.Errorf("unable to feed: %w", err)
	}
	return a, nil
}

func (a *application) feed() error {

	parallelProcessing := []domain.Track{
		domain.NewTrack("Logic Gatekeeper", "", "/media/danimal_cannon_and_zef/parallel_processing/Danimal_Cannon___Zef___Logic_Gatekeeper__uLtRVm_tkxI_/manifest.mpd", time.Minute*5, nil),
		domain.NewTrack("Legacy", "", "/media/danimal_cannon_and_zef/parallel_processing/Danimal_Cannon___Zef___Legacy__pD3JoYx9hfw_/manifest.mpd", time.Minute*5, nil),
		domain.NewTrack("Chronos", "", "/media/danimal_cannon_and_zef/parallel_processing/Danimal_Cannon___Zef___Chronos__eLG_K7X2BaE_/manifest.mpd", time.Minute*5, nil),
		domain.NewTrack("Rhapsody", "", "/media/danimal_cannon_and_zef/parallel_processing/Danimal_Cannon___Zef___Rhapsody__5MSeUK93_Sk_/manifest.mpd", time.Minute*5, nil),
		domain.NewTrack("Glitch", "", "/media/danimal_cannon_and_zef/parallel_processing/Danimal_Cannon___Zef___Glitch__upX7bibwJBw_/manifest.mpd", time.Minute*5, nil),
		domain.NewTrack("The Lunar Whale", "", "/media/danimal_cannon_and_zef/parallel_processing/Danimal_Cannon___Zef___The_Lunar_Whale__p1XpTZPILXM_/manifest.mpd", time.Minute*5, nil),
		domain.NewTrack("Corrupted", "", "/media/danimal_cannon_and_zef/parallel_processing/Danimal_Cannon___Zef___Corrupted__6j_eb_8huQc_/manifest.mpd", time.Minute*5, nil),
		domain.NewTrack("Parallel Processing", "", "/media/danimal_cannon_and_zef/parallel_processing/Danimal_Cannon___Zef___Parallel_Processing__1Fuch5gfUq0_/manifest.mpd", time.Minute*5, nil),
	}

	_, err := a.repo.AddRelease(
		"Parallel Processing",
		domain.ReleaseTypeAlbum,
		"/media/danimal_cannon_and_zef/parallel_processing/parallel_processing_cover.jpg",
		domain.NewDate(2013, 1, 15),
		domain.NewArtist("Danimal Cannon", "/media/danimal_cannon_and_zef/danimal_cannon_profile.jpg"),
		parallelProcessing)
	if err != nil {
		return fmt.Errorf("unable to add release: %w", err)
	}

	liquidLabVol12 := []domain.Track{
		domain.NewTrack("Live set", "", "/media/kream/liquid_lab_vol_12/KREAM_Presents___LIQUID___LAB_Vol__12__Adriatique__Prospa__WhoMadeWho__at_Vemork__LUeHFSkbLlo_/manifest.mpd", time.Minute*5, []domain.Segment{
			domain.NewSegment("KREAM - Liquid Lab Intro", 67),
			domain.NewSegment("Dyzen - She Likes vs", 82),
			domain.NewSegment("RÜFÜS DU SOL - Inhale (SCRIPT Edit) vs", 128),
			domain.NewSegment("RÜFÜS DU SOL - Lately", 189),
			domain.NewSegment("NORRA - This Love", 265),
			domain.NewSegment("RY X - Only vs", 325),
			domain.NewSegment("Simon Doty & DJ Pierre - Come Together", 372),
			domain.NewSegment("Dom Dolla ft Daya - Dreamin’ (KREAM Remix)", 449),
			domain.NewSegment("KREAM & Volaris - ID (Set You Free)", 713),
			domain.NewSegment("Vintage Culture & Tom Breu ft Maverick Sabre - Weak (Acappella) vs", 894),
			domain.NewSegment("Cristoph & Harry Diamond - Hold Me Close vs", 946),
			domain.NewSegment("Florence + The Machine - You've Got The Love (Acappella)", 1022),
			domain.NewSegment("Aaron Hibell ft Felsmann + Tiley - Levitation vs", 1235),
			domain.NewSegment("Elderbrook ft George FitzGerald - Glad I Found You (Acappella) vs", 1280),
			domain.NewSegment("Morgin Madison & Adam Nazar vs Crimsen - Closer (Ryan Lucian Remix)", 1341),
			domain.NewSegment("Prospa ft RAHH - This Rhythm (Acappella) vs", 1418),
			domain.NewSegment("Syence - The Distance", 1448),
			domain.NewSegment("Eric Prydz - 2night vs", 1616),
			domain.NewSegment("Emmit Fenn - The Chase (Acappella)", 1632),
			domain.NewSegment("Adriatique & WhoMadeWho - Miracle (RÜFÜS DU SOL Remix)", 1920),
			domain.NewSegment("KREAM & Ruback - Se Que Quiere", 2073),
			domain.NewSegment("Silar - The Tunnel vs Pa Salieu ft Obongjayar - Style & Fashion (Acappella) vs", 2349),
			domain.NewSegment("Dan Sushi - Orbital", 2418),
			domain.NewSegment("Son Of Son - Lost Control (Acappella) vs", 2494),
			domain.NewSegment("Samm & Ajna - Move vs", 2525),
			domain.NewSegment("PACS - Hyperdrive", 2615),
			domain.NewSegment("Lil Yachty ft Future & Playboi Carti - Flex Up (Acappella) vs", 2693),
			domain.NewSegment("Goom Gum - Staccato", 2738),
			domain.NewSegment("Fideles ft Be No Rain - Night After Night (CamelPhat Remix) vs", 2942),
			domain.NewSegment("Anakim & Dark Heart vs Frýnn - Seconds Away vs", 2998),
			domain.NewSegment("Diplo & HUGEL ft Malou vs Yuna - Forever (Acappella) vs", 3039),
			domain.NewSegment("Adriatique & Solique vs ALSO ASTIR - Changing Colors", 3099),
			domain.NewSegment("Dyzen - Try vs", 3177),
			domain.NewSegment("Cristoph & Pete Tong ft Paul Rogers - Where's The Music Gone", 3221),
			domain.NewSegment("KREAM - Manta", 3390),
			domain.NewSegment("Rivo - Last Night vs", 3696),
			domain.NewSegment("Nils Hoffmann - Closer (Simon Doty Remix)", 3740),
		}),
	}

	_, err = a.repo.AddRelease(
		"Liquid Lab Vol 12",
		domain.ReleaseTypeLiveSet,
		"/media/kream/liquid_lab_vol_12/liquid_lab_vol_12_cover.jpg",
		domain.NewDate(2025, 7, 5),
		domain.NewArtist("Kream", "/media/kream/kream_profile_cover.jpg"),
		liquidLabVol12)
	if err != nil {
		return fmt.Errorf("unable to add release: %w", err)
	}

	gatesOfMetalFriedChickenOfDeath := []domain.Track{
		domain.NewTrack("Away Doom", "", "/media/massacration/gates_of_metal_fried_chicken_of_death/Massacration___Away_Doom__QZurWCQGSkA_/manifest.mpd", time.Minute*5, nil),
		domain.NewTrack("Cereal Metal", "", "/media/massacration/gates_of_metal_fried_chicken_of_death/Massacration___Cereal_Metal__JFP9dUUGtAo_/manifest.mpd", time.Minute*5, nil),
		domain.NewTrack("Evil Papagali", "", "/media/massacration/gates_of_metal_fried_chicken_of_death/Massacration___Evil_Papagali__VTVYAqKdur0_/manifest.mpd", time.Minute*5, nil),
		domain.NewTrack("Feel The Fire From Barbecue", "", "/media/massacration/gates_of_metal_fried_chicken_of_death/Massacration___Feel_The_Fire_From_Barbecue__xH7VsEcZtN4_/manifest.mpd", time.Minute*5, nil),
		domain.NewTrack("Intro", "", "/media/massacration/gates_of_metal_fried_chicken_of_death/Massacration___Intro__dLylOkbplPQ_/manifest.mpd", time.Minute*5, nil),
		domain.NewTrack("Let's Ride To Metal Land The Passage is R$1,0", "", "/media/massacration/gates_of_metal_fried_chicken_of_death/Massacration___Let_s_Ride_To_Metal_Land_The_Passage_is_R_1_0__wGhigjK8H2c_/manifest.mpd", time.Minute*5, nil),
		domain.NewTrack("Metal Bucetation", "", "/media/massacration/gates_of_metal_fried_chicken_of_death/Massacration___Metal_Bucetation__rlft_Ff4cfQ_/manifest.mpd", time.Minute*5, nil),
		domain.NewTrack("Metal Dental Destruction", "", "/media/massacration/gates_of_metal_fried_chicken_of_death/Massacration___Metal_Dental_Destruction__FjIQlPzTcGs_/manifest.mpd", time.Minute*5, nil),
		domain.NewTrack("Metal Glu Glu", "", "/media/massacration/gates_of_metal_fried_chicken_of_death/Massacration___Metal_Glu_Glu__WUtdTBeg11Y_/manifest.mpd", time.Minute*5, nil),
		domain.NewTrack("Metal Is The Law", "", "/media/massacration/gates_of_metal_fried_chicken_of_death/Massacration___Metal_Is_The_Law__VVKxhOosHBg_/manifest.mpd", time.Minute*5, nil),
		domain.NewTrack("Metal Massacre Attack Aruê Aruô", "", "/media/massacration/gates_of_metal_fried_chicken_of_death/Massacration___Metal_Massacre_Attack_Aru__Aru___g6lQl5JjbmI_/manifest.mpd", time.Minute*5, nil),
		domain.NewTrack("Metal Milkshake", "", "/media/massacration/gates_of_metal_fried_chicken_of_death/Massacration___Metal_Milkshake__TMgMVrnOYEw_/manifest.mpd", time.Minute*5, nil),
		domain.NewTrack("The God Master", "", "/media/massacration/gates_of_metal_fried_chicken_of_death/Massacration___The_God_Master__3WQqnb0jqpM_/manifest.mpd", time.Minute*5, nil),
	}

	_, err = a.repo.AddRelease(
		"Gates of Metal Fried Chicken of Death",
		domain.ReleaseTypeAlbum,
		"/media/massacration/gates_of_metal_fried_chicken_of_death/gates_of_metal_fried_chicken_of_death.jpg",
		domain.NewDate(2008, 3, 7),
		domain.NewArtist("Massacration", "/media/massacration/massacration_profile_cover.webp"),
		gatesOfMetalFriedChickenOfDeath)
	if err != nil {
		return fmt.Errorf("unable to add release: %w", err)
	}

	theOphidianTrek := []domain.Track{
		domain.NewTrack("Bleed", "", "/media/meshuggah/the_ophidian_trek/Bleed__Live___0F_zJk2oa_w_/manifest.mpd", time.Minute*5, nil),
		domain.NewTrack("Combustion", "", "/media/meshuggah/the_ophidian_trek/Combustion__Live___LaMUcYtV_3M_/manifest.mpd", time.Minute*5, nil),
		domain.NewTrack("Dancers to a Discordant System", "", "/media/meshuggah/the_ophidian_trek/Dancers_to_a_Discordant_System__Live___AnPJKuRasEU_/manifest.mpd", time.Minute*5, nil),
		domain.NewTrack("Demiurge", "", "/media/meshuggah/the_ophidian_trek/Demiurge__Live___ET6BOPI9mpY_/manifest.mpd", time.Minute*5, nil),
		domain.NewTrack("Do Not Look Down", "", "/media/meshuggah/the_ophidian_trek/Do_Not_Look_Down__Live___PCTYWKqUULE_/manifest.mpd", time.Minute*5, nil),
		domain.NewTrack("I Am Colossus", "", "/media/meshuggah/the_ophidian_trek/I_Am_Colossus__Live___FZrTZLaoP_Y_/manifest.mpd", time.Minute*5, nil),
		domain.NewTrack("Lethargica", "", "/media/meshuggah/the_ophidian_trek/Lethargica__Live___CwxAtsPXozc_/manifest.mpd", time.Minute*5, nil),
		domain.NewTrack("Mirrors / In Death – Is Life / In Death – Is Death", "", "/media/meshuggah/the_ophidian_trek/Mind_s_Mirrors___In_Death___Is_Life___In_Death___Is_Death__Live___FNn0z2lPZqw_/manifest.mpd", time.Minute*5, nil),
		domain.NewTrack("New_Millennium_Cyanide_Christ", "", "/media/meshuggah/the_ophidian_trek/New_Millennium_Cyanide_Christ__Live___EKWQvz0OYtY_/manifest.mpd", time.Minute*5, nil),
		domain.NewTrack("Obzen", "", "/media/meshuggah/the_ophidian_trek/Obzen__Live___spZ2GT3kM1U_/manifest.mpd", time.Minute*5, nil),
		domain.NewTrack("Rational Gaze", "", "/media/meshuggah/the_ophidian_trek/Rational_Gaze__Live___rV_GzVzdWxk_/manifest.mpd", time.Minute*5, nil),
		domain.NewTrack("Swarm", "", "/media/meshuggah/the_ophidian_trek/Swarm__Live___FaOdBys4pgQ_/manifest.mpd", time.Minute*5, nil),
		domain.NewTrack("Swarmer", "", "/media/meshuggah/the_ophidian_trek/Swarmer__Live___1HhNYMD1_Zk_/manifest.mpd", time.Minute*5, nil),
		domain.NewTrack("The Hurt That Finds You First", "", "/media/meshuggah/the_ophidian_trek/The_Hurt_That_Finds_You_First__Live___xm7jBTRk4p4_/manifest.mpd", time.Minute*5, nil),
		domain.NewTrack("The Last Vigil", "", "/media/meshuggah/the_ophidian_trek/The_Last_Vigil__Live___FDWRJQuINnk_/manifest.mpd", time.Minute*5, nil),
	}

	_, err = a.repo.AddRelease(
		"The Ophidian Trek (Live)",
		domain.ReleaseTypeAlbum,
		"/media/meshuggah/the_ophidian_trek/ophidian_trek_cover.jpg",
		domain.NewDate(2014, 9, 29),
		domain.NewArtist("Meshuggah", "/media/meshuggah/meshuggah_profile_cover.jpg"),
		theOphidianTrek)
	if err != nil {
		return fmt.Errorf("unable to add release: %w", err)
	}

	obzen := []domain.Track{
		domain.NewTrack("Bleed", "", "/media/meshuggah/obzen/Bleed__GAulPs96ass_/manifest.mpd", time.Minute*5, nil),
		domain.NewTrack("Combustion", "", "/media/meshuggah/obzen/Combustion__RL7RFrInBww_/manifest.mpd", time.Minute*5, nil),
		domain.NewTrack("Dancers_to_a_Discordant_System", "", "/media/meshuggah/obzen/Dancers_to_a_Discordant_System__0DxI_ZOGbXo_/manifest.mpd", time.Minute*5, nil),
		domain.NewTrack("Electric_Red", "", "/media/meshuggah/obzen/Electric_Red__UpTmMSXm9rw_/manifest.mpd", time.Minute*5, nil),
		domain.NewTrack("Lethargica", "", "/media/meshuggah/obzen/Lethargica__Qywg2ZjdqMo_/manifest.mpd", time.Minute*5, nil),
		domain.NewTrack("Obzen", "", "/media/meshuggah/obzen/Obzen__Fc7aeQGLccI_/manifest.mpd", time.Minute*5, nil),
		domain.NewTrack("Pineal_Gland_Optics", "", "/media/meshuggah/obzen/Pineal_Gland_Optics__VgD2Ks_gxxw_/manifest.mpd", time.Minute*5, nil),
		domain.NewTrack("Pravus", "", "/media/meshuggah/obzen/Pravus__6KJ2RCm_bcs_/manifest.mpd", time.Minute*5, nil),
		domain.NewTrack("This_Spiteful_Snake", "", "/media/meshuggah/obzen/This_Spiteful_Snake__cjjOW5X8nIw_/manifest.mpd", time.Minute*5, nil),
	}
	_, err = a.repo.AddRelease(
		"Obzen",
		domain.ReleaseTypeAlbum,
		"/media/meshuggah/obzen/obsen_cover.jpg",
		domain.NewDate(2008, 3, 7),
		domain.NewArtist("Meshuggah", "/media/meshuggah/meshuggah_profile_cover.jpg"),
		obzen)
	if err != nil {
		return fmt.Errorf("unable to add release: %w", err)
	}

	nothingRemastered := []domain.Track{
		domain.NewTrack("Closed Eye Visuals", "", "/media/meshuggah/nothing_remastered/Closed_Eye_Visuals__Remastered_2006___kZ84SjwKe80_/manifest.mpd", time.Minute*5, nil),
		domain.NewTrack("Glints Collide", "", "/media/meshuggah/nothing_remastered/Glints_Collide__Remastered_2006___rU6oJ0Kdews_/manifest.mpd", time.Minute*5, nil),
		domain.NewTrack("Nebulous", "", "/media/meshuggah/nothing_remastered/Nebulous__Remastered_2006___K6ncYs9Bji8_/manifest.mpd", time.Minute*5, nil),
		domain.NewTrack("Obsidian", "", "/media/meshuggah/nothing_remastered/Obsidian__Remastered_2006___TOkPHeGuWuQ_/manifest.mpd", time.Minute*5, nil),
		domain.NewTrack("Organic Shadows", "", "/media/meshuggah/nothing_remastered/Organic_Shadows__Remastered_2006___JDlJB_ClTdg_/manifest.mpd", time.Minute*5, nil),
		domain.NewTrack("Perpetual Black Second", "", "/media/meshuggah/nothing_remastered/Perpetual_Black_Second__Remastered_2006____NkPYAU1I1U_/manifest.mpd", time.Minute*5, nil),
		domain.NewTrack("Rational Gaze", "", "/media/meshuggah/nothing_remastered/Rational_Gaze__Remastered_2006___IcdlMBvMCjs_/manifest.mpd", time.Minute*5, nil),
		domain.NewTrack("Spasm", "", "/media/meshuggah/nothing_remastered/Spasm__Remastered_2006___aZAYKpYkHNI_/manifest.mpd", time.Minute*5, nil),
		domain.NewTrack("Stengah", "", "/media/meshuggah/nothing_remastered/Stengah__Remastered_2006___ntISKKjf0gk_/manifest.mpd", time.Minute*5, nil),
		domain.NewTrack("Straws Pulled at Random", "", "/media/meshuggah/nothing_remastered/Straws_Pulled_at_Random__Remastered_2006___xA_qx8ht2j0_/manifest.mpd", time.Minute*5, nil),
	}
	_, err = a.repo.AddRelease(
		"Nothing (Remastered 2006)",
		domain.ReleaseTypeAlbum,
		"/media/meshuggah/nothing_remastered/nothing_cover.jpg",
		domain.NewDate(2006, 10, 31),
		domain.NewArtist("Meshuggah", "/media/meshuggah/meshuggah_profile_cover.jpg"),
		nothingRemastered)
	if err != nil {
		return fmt.Errorf("unable to add release: %w", err)
	}

	ridingTheLigthningSongs := []domain.Track{
		domain.NewTrack("Creeping Death", "", "/media/metallica/riding_the_lightning/Creeping_Death__Studio_Version___2h4iqDSzVv0_/manifest.mpd", time.Minute*5, nil),
		domain.NewTrack("Escape", "", "/media/metallica/riding_the_lightning/Escape__Studio_Version___MGdKPy98Byg_/manifest.mpd", time.Minute*5, nil),
		domain.NewTrack("Fade To Black", "", "/media/metallica/riding_the_lightning/Fade_To_Black__Studio_Version___eDZPLSexHVM_/manifest.mpd", time.Minute*5, nil),
		domain.NewTrack("Fight Fire With Fire", "", "/media/metallica/riding_the_lightning/Fight_Fire_With_Fire__Studio_Version___ZnCFWlso_UQ_/manifest.mpd", time.Minute*5, nil),
		domain.NewTrack("For Whom The Bell Tolls", "", "/media/metallica/riding_the_lightning/For_Whom_The_Bell_Tolls__Studio_Version____b6tJMD34qw_/manifest.mpd", time.Minute*5, nil),
		domain.NewTrack("Ride The Lightning", "", "/media/metallica/riding_the_lightning/Ride_The_Lightning__Studio_Version___ArgdUZKslPw_/manifest.mpd", time.Minute*5, nil),
		domain.NewTrack("The Call Of Ktulu", "", "/media/metallica/riding_the_lightning/The_Call_Of_Ktulu__Studio_Version___Z_wccw663BE_/manifest.mpd", time.Minute*5, nil),
		domain.NewTrack("Trapped Under Ice", "", "/media/metallica/riding_the_lightning/Trapped_Under_Ice__Studio_Version___6mLDoLWJKZw_/manifest.mpd", time.Minute*5, nil),
	}

	_, err = a.repo.AddRelease(
		"Riding the Lightning",
		domain.ReleaseTypeAlbum,
		"/media/metallica/riding_the_lightning/riding_the_lightning_cover.jpg",
		domain.NewDate(1984, 7, 27),
		domain.NewArtist("Metallica", "/media/metallica/metallica_profile_cover.jpg"),
		ridingTheLigthningSongs)
	if err != nil {
		return fmt.Errorf("unable to add release: %w", err)
	}

	masterOfPuppetsSongs := []domain.Track{
		domain.NewTrack("Battery", "", "/media/metallica/master_of_puppets/Battery__Remastered___uzlOcupu5UE_/manifest.mpd", time.Minute*5, nil),
		domain.NewTrack("Damage Inc.", "", "/media/metallica/master_of_puppets/Damage__Inc___Remastered___Abe3AZhcGQs_/manifest.mpd", time.Minute*5, nil),
		domain.NewTrack("Disposable Heroes", "", "/media/metallica/master_of_puppets/Disposable_Heroes__Remastered___p3Y8VSVyYN8_/manifest.mpd", time.Minute*5, nil),
		domain.NewTrack("Leper Messiah", "", "/media/metallica/master_of_puppets/Leper_Messiah__Remastered___dJp5r4HdRn4_/manifest.mpd", time.Minute*5, nil),
		domain.NewTrack("Master Of Puppets", "", "/media/metallica/master_of_puppets/Master_Of_Puppets__Remastered___u6LahTuw02c_/manifest.mpd", time.Minute*5, nil),
		domain.NewTrack("Orion", "", "/media/metallica/master_of_puppets/Orion__Remastered___z7bUJPj_8v0_/manifest.mpd", time.Minute*5, nil),
		domain.NewTrack("The Thing That Should Never Be", "", "/media/metallica/master_of_puppets/The_Thing_That_Should_Not_Be__Remastered___gm9c_QpuMms_/manifest.mpd", time.Minute*5, nil),
		domain.NewTrack("Welcome Home (Sanitarium)", "", "/media/metallica/master_of_puppets/Welcome_Home__Sanitarium___Remastered___G_868UwoJvM_/manifest.mpd", time.Minute*5, nil),
	}

	id, err := a.repo.AddRelease(
		"Master of Puppets",
		domain.ReleaseTypeAlbum,
		"/media/metallica/master_of_puppets/master_of_puppets_cover.jpg",
		domain.NewDate(1986, 3, 3),
		domain.NewArtist("Metallica", "/media/metallica/metallica_profile_cover.jpg"),
		masterOfPuppetsSongs)
	if err != nil {
		return fmt.Errorf("unable to add release: %w", err)
	}
	fmt.Println("releaseID " + id.String())
	return nil
}
