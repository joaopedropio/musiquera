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

	gatesOfMetalFriedChickenOfDeath := []domain.Song{
		domain.NewSong("Away Doom", "", "/media/massacration/gates_of_metal_fried_chicken_of_death/Massacration___Away_Doom__QZurWCQGSkA_/manifest.mpd", time.Minute*5),
		domain.NewSong("Cereal Metal", "", "/media/massacration/gates_of_metal_fried_chicken_of_death/Massacration___Cereal_Metal__JFP9dUUGtAo_/manifest.mpd", time.Minute*5),
		domain.NewSong("Evil Papagali", "", "/media/massacration/gates_of_metal_fried_chicken_of_death/Massacration___Evil_Papagali__VTVYAqKdur0_/manifest.mpd", time.Minute*5),
		domain.NewSong("Feel The Fire From Barbecue", "", "/media/massacration/gates_of_metal_fried_chicken_of_death/Massacration___Feel_The_Fire_From_Barbecue__xH7VsEcZtN4_/manifest.mpd", time.Minute*5),
		domain.NewSong("Intro", "", "/media/massacration/gates_of_metal_fried_chicken_of_death/Massacration___Intro__dLylOkbplPQ_/manifest.mpd", time.Minute*5),
		domain.NewSong("Let's Ride To Metal Land The Passage is R$1,0", "", "/media/massacration/gates_of_metal_fried_chicken_of_death/Massacration___Let_s_Ride_To_Metal_Land_The_Passage_is_R_1_0__wGhigjK8H2c_/manifest.mpd", time.Minute*5),
		domain.NewSong("Metal Bucetation", "", "/media/massacration/gates_of_metal_fried_chicken_of_death/Massacration___Metal_Bucetation__rlft_Ff4cfQ_/manifest.mpd", time.Minute*5),
		domain.NewSong("Metal Dental Destruction", "", "/media/massacration/gates_of_metal_fried_chicken_of_death/Massacration___Metal_Dental_Destruction__FjIQlPzTcGs_/manifest.mpd", time.Minute*5),
		domain.NewSong("Metal Glu Glu", "", "/media/massacration/gates_of_metal_fried_chicken_of_death/Massacration___Metal_Glu_Glu__WUtdTBeg11Y_/manifest.mpd", time.Minute*5),
		domain.NewSong("Metal Is The Law", "", "/media/massacration/gates_of_metal_fried_chicken_of_death/Massacration___Metal_Is_The_Law__VVKxhOosHBg_/manifest.mpd", time.Minute*5),
		domain.NewSong("Metal Massacre Attack Aruê Aruô", "", "/media/massacration/gates_of_metal_fried_chicken_of_death/Massacration___Metal_Massacre_Attack_Aru__Aru___g6lQl5JjbmI_/manifest.mpd", time.Minute*5),
		domain.NewSong("Metal Milkshake", "", "/media/massacration/gates_of_metal_fried_chicken_of_death/Massacration___Metal_Milkshake__TMgMVrnOYEw_/manifest.mpd", time.Minute*5),
		domain.NewSong("The God Master", "", "/media/massacration/gates_of_metal_fried_chicken_of_death/Massacration___The_God_Master__3WQqnb0jqpM_/manifest.mpd", time.Minute*5),
	}

	_, err := a.repo.AddAlbum(
		"Gates of Metal Fried Chicken of Death",
		domain.NewDate(2008, 3, 7),
		domain.NewArtist("Massacration"),
		gatesOfMetalFriedChickenOfDeath)
	if err != nil {
		return fmt.Errorf("unable to add album: %w", err)
	}

	theOphidianTrek := []domain.Song{
		domain.NewSong("Bleed", "", "/media/meshuggah/the_ophidian_trek/Bleed__Live___0F_zJk2oa_w_/manifest.mpd", time.Minute*5),
		domain.NewSong("Combustion", "", "/media/meshuggah/the_ophidian_trek/Combustion__Live___LaMUcYtV_3M_/manifest.mpd", time.Minute*5),
		domain.NewSong("Dancers to a Discordant System", "", "/media/meshuggah/the_ophidian_trek/Dancers_to_a_Discordant_System__Live___AnPJKuRasEU_/manifest.mpd", time.Minute*5),
		domain.NewSong("Demiurge", "", "/media/meshuggah/the_ophidian_trek/Demiurge__Live___ET6BOPI9mpY_/manifest.mpd", time.Minute*5),
		domain.NewSong("Do Not Look Down", "", "/media/meshuggah/the_ophidian_trek/Do_Not_Look_Down__Live___PCTYWKqUULE_/manifest.mpd", time.Minute*5),
		domain.NewSong("I Am Colossus", "", "/media/meshuggah/the_ophidian_trek/I_Am_Colossus__Live___FZrTZLaoP_Y_/manifest.mpd", time.Minute*5),
		domain.NewSong("Lethargica", "", "/media/meshuggah/the_ophidian_trek/Lethargica__Live___CwxAtsPXozc_/manifest.mpd", time.Minute*5),
		domain.NewSong("Mirrors / In Death – Is Life / In Death – Is Death", "", "/media/meshuggah/the_ophidian_trek/Mind_s_Mirrors___In_Death___Is_Life___In_Death___Is_Death__Live___FNn0z2lPZqw_/manifest.mpd", time.Minute*5),
		domain.NewSong("New_Millennium_Cyanide_Christ", "", "/media/meshuggah/the_ophidian_trek/New_Millennium_Cyanide_Christ__Live___EKWQvz0OYtY_/manifest.mpd", time.Minute*5),
		domain.NewSong("Obzen", "", "/media/meshuggah/the_ophidian_trek/Obzen__Live___spZ2GT3kM1U_/manifest.mpd", time.Minute*5),
		domain.NewSong("Rational Gaze", "", "/media/meshuggah/the_ophidian_trek/Rational_Gaze__Live___rV_GzVzdWxk_/manifest.mpd", time.Minute*5),
		domain.NewSong("Swarm", "", "/media/meshuggah/the_ophidian_trek/Swarm__Live___FaOdBys4pgQ_/manifest.mpd", time.Minute*5),
		domain.NewSong("Swarmer", "", "/media/meshuggah/the_ophidian_trek/Swarmer__Live___1HhNYMD1_Zk_/manifest.mpd", time.Minute*5),
		domain.NewSong("The Hurt That Finds You First", "", "/media/meshuggah/the_ophidian_trek/The_Hurt_That_Finds_You_First__Live___xm7jBTRk4p4_/manifest.mpd", time.Minute*5),
		domain.NewSong("The Last Vigil", "", "/media/meshuggah/the_ophidian_trek/The_Last_Vigil__Live___FDWRJQuINnk_/manifest.mpd", time.Minute*5),
	}

	_, err = a.repo.AddAlbum(
		"The Ophidian Trek (Live)",
		domain.NewDate(2014, 9, 29),
		domain.NewArtist("Meshuggah"),
		theOphidianTrek)
	if err != nil {
		return fmt.Errorf("unable to add album: %w", err)
	}

	obzen := []domain.Song{
		domain.NewSong("Bleed", "", "/media/meshuggah/obzen/Bleed__GAulPs96ass_/manifest.mpd", time.Minute*5),
		domain.NewSong("Combustion", "", "/media/meshuggah/obzen/Combustion__RL7RFrInBww_/manifest.mpd", time.Minute*5),
		domain.NewSong("Dancers_to_a_Discordant_System", "", "/media/meshuggah/obzen/Dancers_to_a_Discordant_System__0DxI_ZOGbXo_/manifest.mpd", time.Minute*5),
		domain.NewSong("Electric_Red", "", "/media/meshuggah/obzen/Electric_Red__UpTmMSXm9rw_/manifest.mpd", time.Minute*5),
		domain.NewSong("Lethargica", "", "/media/meshuggah/obzen/Lethargica__Qywg2ZjdqMo_/manifest.mpd", time.Minute*5),
		domain.NewSong("Obzen", "", "/media/meshuggah/obzen/Obzen__Fc7aeQGLccI_/manifest.mpd", time.Minute*5),
		domain.NewSong("Pineal_Gland_Optics", "", "/media/meshuggah/obzen/Pineal_Gland_Optics__VgD2Ks_gxxw_/manifest.mpd", time.Minute*5),
		domain.NewSong("Pravus", "", "/media/meshuggah/obzen/Pravus__6KJ2RCm_bcs_/manifest.mpd", time.Minute*5),
		domain.NewSong("This_Spiteful_Snake", "", "/media/meshuggah/obzen/This_Spiteful_Snake__cjjOW5X8nIw_/manifest.mpd", time.Minute*5),
	}
	_, err = a.repo.AddAlbum(
		"Obzen",
		domain.NewDate(2008, 3, 7),
		domain.NewArtist("Meshuggah"),
		obzen)
	if err != nil {
		return fmt.Errorf("unable to add album: %w", err)
	}

	nothingRemastered := []domain.Song{
		domain.NewSong("Closed Eye Visuals", "", "/media/meshuggah/nothing_remastered/Closed_Eye_Visuals__Remastered_2006___kZ84SjwKe80_/manifest.mpd", time.Minute*5),
		domain.NewSong("Glints Collide", "", "/media/meshuggah/nothing_remastered/Glints_Collide__Remastered_2006___rU6oJ0Kdews_/manifest.mpd", time.Minute*5),
		domain.NewSong("Nebulous", "", "/media/meshuggah/nothing_remastered/Nebulous__Remastered_2006___K6ncYs9Bji8_/manifest.mpd", time.Minute*5),
		domain.NewSong("Obsidian", "", "/media/meshuggah/nothing_remastered/Obsidian__Remastered_2006___TOkPHeGuWuQ_/manifest.mpd", time.Minute*5),
		domain.NewSong("Organic Shadows", "", "/media/meshuggah/nothing_remastered/Organic_Shadows__Remastered_2006___JDlJB_ClTdg_/manifest.mpd", time.Minute*5),
		domain.NewSong("Perpetual Black Second", "", "/media/meshuggah/nothing_remastered/Perpetual_Black_Second__Remastered_2006____NkPYAU1I1U_/manifest.mpd", time.Minute*5),
		domain.NewSong("Rational Gaze", "", "/media/meshuggah/nothing_remastered/Rational_Gaze__Remastered_2006___IcdlMBvMCjs_/manifest.mpd", time.Minute*5),
		domain.NewSong("Spasm", "", "/media/meshuggah/nothing_remastered/Spasm__Remastered_2006___aZAYKpYkHNI_/manifest.mpd", time.Minute*5),
		domain.NewSong("Stengah", "", "/media/meshuggah/nothing_remastered/Stengah__Remastered_2006___ntISKKjf0gk_/manifest.mpd", time.Minute*5),
		domain.NewSong("Straws Pulled at Random", "", "/media/meshuggah/nothing_remastered/Straws_Pulled_at_Random__Remastered_2006___xA_qx8ht2j0_/manifest.mpd", time.Minute*5),
	}
	_, err = a.repo.AddAlbum(
		"Nothing (Remastered 2006)",
		domain.NewDate(2006, 10, 31),
		domain.NewArtist("Meshuggah"),
		nothingRemastered)
	if err != nil {
		return fmt.Errorf("unable to add album: %w", err)
	}

	ridingTheLigthningSongs := []domain.Song{
		domain.NewSong("Creeping Death", "", "/media/metallica/riding_the_lightning/Creeping_Death__Studio_Version___2h4iqDSzVv0_/manifest.mpd", time.Minute*5),
		domain.NewSong("Escape", "", "/media/metallica/riding_the_lightning/Escape__Studio_Version___MGdKPy98Byg_/manifest.mpd", time.Minute*5),
		domain.NewSong("Fade To Black", "", "/media/metallica/riding_the_lightning/Fade_To_Black__Studio_Version___eDZPLSexHVM_/manifest.mpd", time.Minute*5),
		domain.NewSong("Fight Fire With Fire", "", "/media/metallica/riding_the_lightning/Fight_Fire_With_Fire__Studio_Version___ZnCFWlso_UQ_/manifest.mpd", time.Minute*5),
		domain.NewSong("For Whom The Bell Tolls", "", "/media/metallica/riding_the_lightning/For_Whom_The_Bell_Tolls__Studio_Version____b6tJMD34qw_/manifest.mpd", time.Minute*5),
		domain.NewSong("Ride The Lightning", "", "/media/metallica/riding_the_lightning/Ride_The_Lightning__Studio_Version___ArgdUZKslPw_/manifest.mpd", time.Minute*5),
		domain.NewSong("The Call Of Ktulu", "", "/media/metallica/riding_the_lightning/The_Call_Of_Ktulu__Studio_Version___Z_wccw663BE_/manifest.mpd", time.Minute*5),
		domain.NewSong("Trapped Under Ice", "", "/media/metallica/riding_the_lightning/Trapped_Under_Ice__Studio_Version___6mLDoLWJKZw_/manifest.mpd", time.Minute*5),
	}

	_, err = a.repo.AddAlbum(
		"Riding the Lightning",
		domain.NewDate(1984, 7, 27),
		domain.NewArtist("Metallica"),
		ridingTheLigthningSongs)
	if err != nil {
		return fmt.Errorf("unable to add album: %w", err)
	}

	masterOfPuppetsSongs := []domain.Song{
		domain.NewSong("Battery", "", "/media/metallica/master_of_puppets/Battery__Remastered___uzlOcupu5UE_/manifest.mpd", time.Minute*5),
		domain.NewSong("Damage Inc.", "", "/media/metallica/master_of_puppets/Damage__Inc___Remastered___Abe3AZhcGQs_/manifest.mpd", time.Minute*5),
		domain.NewSong("Disposable Heroes", "", "/media/metallica/master_of_puppets/Disposable_Heroes__Remastered___p3Y8VSVyYN8_/manifest.mpd", time.Minute*5),
		domain.NewSong("Leper Messiah", "", "/media/metallica/master_of_puppets/Leper_Messiah__Remastered___dJp5r4HdRn4_/manifest.mpd", time.Minute*5),
		domain.NewSong("Master Of Puppets", "", "/media/metallica/master_of_puppets/Master_Of_Puppets__Remastered___u6LahTuw02c_/manifest.mpd", time.Minute*5),
		domain.NewSong("Orion", "", "/media/metallica/master_of_puppets/Orion__Remastered___z7bUJPj_8v0_/manifest.mpd", time.Minute*5),
		domain.NewSong("The Thing That Should Never Be", "", "/media/metallica/master_of_puppets/The_Thing_That_Should_Not_Be__Remastered___gm9c_QpuMms_/manifest.mpd", time.Minute*5),
		domain.NewSong("Welcome Home (Sanitarium)", "", "/media/metallica/master_of_puppets/Welcome_Home__Sanitarium___Remastered___G_868UwoJvM_/manifest.mpd", time.Minute*5),
	}

	id, err := a.repo.AddAlbum(
		"Master of Puppets",
		domain.NewDate(1986, 3, 3),
		domain.NewArtist("Metallica"),
		masterOfPuppetsSongs)
	if err != nil {
		return fmt.Errorf("unable to add album: %w", err)
	}
	fmt.Println("albumID " + id.String())
	return nil
}
