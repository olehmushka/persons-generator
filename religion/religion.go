package religion

import (
	"fmt"

	"persons_generator/entities"
)

func New() *entities.Religion {
	fmt.Println("[religion.New] started")
	defer fmt.Println("[religion.New] finished")

	r := &entities.Religion{}
	generateDoctrine(r)
	generateTheology(r)
	generateTaboos(r)
	generateClerics(r)

	return r
}

func NewMany(n int) []*entities.Religion {
	religions := make([]*entities.Religion, 0, n)
	for i := 0; i < n; i++ {
		religions = append(religions, New())
	}

	return religions
}

// temples, shelters, holy scripture, symbolics, initiation
// religious art, religious literature, reliquary, sacred places
// preachers, forced conversion, masters, missionary outreach

// monotheism

// rich ideology: cult property, indulgences, tithe, primacy of the pope, tax to support the poor
// poor ideology: holiness, celibacy, flagellantism

// developed ideology: seminaries, libraries, feed the world, patronship
// undeveloped ideology: secret sect, sacrificial offesring, rites and chants

// militant ideology: holy army, warrior monks, defenders of Faith
// peaceful ideology: non-violent resistance, swords to plowshares, no more killing, world religion

// primitive ideology: great fast,, repentance, martyrdom
// organized ideology: confession, commandments, great afterlife, devotional code

// divine law, sacrament, justice of heaven

// polytheism

// rich ideology: buying your sins, inviolacy of temples, god of king cult, cult of justice, god of prosperity and wealth, sacred stone, patron saint of merchants
// poor ideology: oracles, diviners, libations, holy animals, all father's watchmen, ordeals,

// developed ideology: seminaries, libraries, schools of philosophers, kalokagathos, galdbarok, runes, volhvs, periapts
// undeveloped ideology: hecatombs, pharmakos, inferiae, god of erebus, trickster, samhain

// militant ideology: god of war cult, gymnasium schools, holy army. marching hymns, agoge, berserkers, heavenly palace
// peaceful ideology: temple healers, sport contest, oath of healer, veneration of the ancestors, blood oath, truce

// primitive ideology: symposium, music, theatre,, liturgical drama, mead of poetry, visas, sages
// organized ideology: theogony, scapegoat, psychopomp, spoken word, norns, tapestry, prophesy, honorable death

// mother goddess, deed of expiation, mysteries, nithing

// deism

// rich ideology: emperor's doctrine
// poor ideology: animal messengers

// developed ideology: seminaries, calligraphy, sanctuary of mind
// undeveloped ideology: phantoms, living pillars, demonic deal

// militant ideology: armor of god, warrior's path
// peaceful ideology:

// primitive ideology:
// organized ideology: calendar of feasts, prayer spells, spirits of death

// mono no aware, cleansing rituals
