package db

import (
	"database/sql"
	"go-trigger-app/model"
)

// DB interface with the new method
type DB interface {
	CreateUser(user *model.User, passwordHash string) (int, error)
	GetUserByUsername(username string) (*model.User, string, error)
	GetTechnologiesByUserID(userID int) ([]*model.Technology, error)
	AddDefaultLeaksForUser(userID int) error
}

type PostgresDB struct {
	db *sql.DB
}

func NewDB(db *sql.DB) DB {
	return PostgresDB{db: db}
}

// ... (CreateUser, GetUserByUsername, GetTechnologiesByUserID functions remain the same) ...

func (d PostgresDB) CreateUser(user *model.User, passwordHash string) (int, error) {
	var id int
	err := d.db.QueryRow("INSERT INTO users (username, password_hash) VALUES ($1, $2) RETURNING id", user.Username, passwordHash).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (d PostgresDB) GetUserByUsername(username string) (*model.User, string, error) {
	user := new(model.User)
	var passwordHash string
	err := d.db.QueryRow("SELECT id, username, password_hash FROM users WHERE username = $1", username).Scan(&user.ID, &user.Username, &passwordHash)
	if err != nil {
		return nil, "", err
	}
	return user, passwordHash, nil
}

func (d PostgresDB) GetTechnologiesByUserID(userID int) ([]*model.Technology, error) {
	rows, err := d.db.Query("select id, user_id, name, details from technologies where user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var tech []*model.Technology
	for rows.Next() {
		t := new(model.Technology)
		err = rows.Scan(&t.ID, &t.UserID, &t.Name, &t.Details)
		if err != nil {
			return nil, err
		}
		tech = append(tech, t)
	}
	return tech, nil
}

// New function to add all the default leaks for a new user
func (d PostgresDB) AddDefaultLeaksForUser(userID int) error {
	leaks := []model.Technology{
		{Name: "Modern Business Solutions", Details: "In October 2016, a large Mongo database file containing tens of millions of accounts was published by Twitter. The database contained more than 58 million unique emails, as well as IP addresses, names, home addresses, gender, positions, dates of birth and telephone numbers. The data is attributed to Modern Business Solutions, a company that provides data storage and database hosting solutions."},
		{Name: "Twitter (Partial)", Details: "In January 2022, a vulnerability in the Twitter platform allowed the creation of a database of emails and phone numbers of millions of users. Twitter said the vulnerability stems from a bug introduced in June 2021. The data included publicly available information: username, display name, biography, location, and profile photo. 6.7 million email addresses were affected."},
		{Name: "Twitter 200M", Details: "In January 2022, a vulnerability in the Twitter platform allowed a hacker to create a database of emails and phone numbers of 200 million users of the social platform. In August 2022, Twitter reported that the vulnerability was related to a bug introduced in June 2021. The data included emails, phone numbers, username, biography, location and profile photo."},
		{Name: "Cloudata", Details: "Large collection of email-pass data. The database was collected from many files on May 18, 2023. Initially, all databases weighed 338 GB (11 billion rows). After removing duplicates and data from the collections, about 2 billion remained."},
		{Name: "BreachedForum Combolists", Details: "Breached.vc is the most famous site for publishing various leaks. This database contains all files from the “Combolists” section, downloaded on 02/22/2023. This section had 2200 databases containing 170GB of data (5 billion rows). After deduplication and removal of rows from the collections, only 770 million remained. The data contains only emails and passwords."},
		{Name: "GameTuts", Details: "Apparently in early 2015, video game site GameTuts suffered a data breach that exposed over 2 million user accounts. The exposed data included usernames, email addresses and IP addresses, as well as salted MD5 hashes of passwords."},
		{Name: "USA Voters", Details: "This collection of files contains voter databases for 27 US states. Much of this data was purchased by Omnipotent for RaidForums in 2018 (most states charge money to download voter data). The databases contain full names, addresses, telephones, emails and parties selected during voting."},
		{Name: "StockX", Details: "In July 2019, fashion and sneaker trading platform StockX suffered a data breach that was subsequently sold on a darknet marketplace. The data included 6.8 million unique email addresses, as well as names, physical purchase addresses, and salted MD5 hashes of passwords."},
		{Name: "NexusMods", Details: "In December 2015, game modding site Nexus Mods released a statement notifying users that they had been hacked. They subsequently dated the hack to July 2013, although there is evidence that the data was sold several months earlier. The hack contained usernames, email addresses and passwords as salted hashes."},
		{Name: "Ga$$Pacc", Details: "In 2020, a set of decrypted databases containing 580 million emails and passwords was made publicly available. The data was obtained from 243 different leaks and contained only emails and unencrypted passwords."},
		{Name: "Kixify", Details: "It is an international online platform for buying and selling sneakers and sports shoes. It was hacked in 2022, which led to the leak of 3 million records with emails, names and passwords in the form of MD5 hashes without salt. The site was later hacked again in September 2023, causing another 800 thousand records to leak."},
		{Name: "Collection #1", Details: "In January 2019, 5 collections appeared on a popular hacker forum, representing a combination of data from many other leaks. This is the first collection found. It contained 12,300 files with a total size of 87 GB. The collection contains more than 2.5 billion lines with emails and passwords for them. There were only 1.25 billion unique rows. For many strings, the site where the leak occurred was indicated."},
		{Name: "Collection #5", Details: "In January 2019, 5 collections appeared on a popular hacker forum, representing a combination of data from many other leaks. The fifth collection contains 16 thousand files with a total size of 42 gigabytes. The files contained 1.25 billion pairs of emails and passwords, but there were only 600 million unique strings. Unlike other collections, there was often information about which site the data was leaked from."},
		{Name: "Mathway", Details: "In January 2020, math problem-solving website Mathway suffered a data breach. As a result, more than 25 million records were exposed. The data was subsequently sold on the darknet market and included names, Google and Facebook IDs, email addresses, and base64-encoded salted MD5 password hashes."},
		{Name: "Collection", Details: "In January 2019, a collection of 5 parts appeared on a hacker forum, representing a combination of data from different sites. The collection contained 108 thousand files with a total size of 870GB. These files contained 26 billion lines containing email and password. However, many values ​​were repeated many times, and there were only 4.3 billion unique pairs. And yet this collection is considered the biggest leak in history."},
		{Name: "Collection of 7,651 databases", Details: "In December 2022, all databases from the hacker data trading site were made publicly available. A total of 7,651 leaks were published. The data contained only plaintext emails and passwords. A total of 493 million records in all databases."},
		{Name: "Chegg", Details: "In April 2018, textbook rental service Chegg suffered a leak that affected 40 million subscribers. The data included email, usernames, and passwords in the form of unsalted MD5 hashes."},
		{Name: "Poshmark", Details: "In mid-2018, social commerce marketplace Poshmark suffered a data breach that exposed the accounts of 36 million users. The data included email addresses, usernames, gender, location, and passwords in the form of bcrypt hashes."},
		{Name: "Edmodo", Details: "In May 2017, the educational platform Edmodo was hacked, exposing 77 million records containing 43 million emails. Usernames and bcrypt password hashes were also published."},
		{Name: "TeraBase64", Details: "A huge collection of files published in February 2020 by a person with the handle @HTTSMVKCOM. It contained 3.2 billion lines of plain text emails and passwords, but only 1.28 billion unique lines. All this data was most likely obtained from many other leaks."},
		{Name: "RobinHood", Details: "In November 2021, online trading platform Robinhood suffered a data breach. Hackers used social engineering on a customer service representative. The incident exposed more than 5 million email addresses and 2 million customer names."},
		{Name: "Zeeroq", Details: "Around 2019, about 250 combo players were found in the open directory on the demo.zeeroq.com domain. The number of entries was more than 200 million. The leak contained emails and passwords in plain text."},
		{Name: "MyHeritage", Details: "In October 2017, the genealogy site MyHeritage was leaked. This became known 7 months later. More than 92 million customer records were exposed, including email and password hashes with SHA-1 salt."},
	}

	for _, leak := range leaks {
		_, err := d.db.Exec("INSERT INTO technologies (user_id, name, details) VALUES ($1, $2, $3)", userID, leak.Name, leak.Details)
		if err != nil {
			return err
		}
	}
	return nil
}
