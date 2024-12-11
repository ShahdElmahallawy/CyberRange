package db

import (
	"os"
)

func AddChallenge(name, description, difficulty, flag, points, hint, category, attachments string) error {

	_, err := DB.Exec("INSERT INTO ctf_challenges (name, description, difficulty, flag, points, hint, category, attachments) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)", name, description, difficulty, flag, points, hint, category, attachments)
	if err != nil {

		return err
	}

	return nil

}

func DeleteChallenge(name string) error {
	_, err := DB.Exec("DELETE FROM ctf_challenges WHERE name = $1", name)
	if err != nil {
		return err
	}

	path := "CyberRange/CTF/" + name
	err = os.RemoveAll(path)
	if err != nil {
		return err
	}

	return nil

}

func GetSettings(ctfType string) (map[string]string, error) {

	var query string

	if ctfType == "jeopardy" {
		query = "SELECT flag, status,set_for_release,release_date FROM ctf WHERE id = 1"
	} else if ctfType == "attack-defense" {
		query = "SELECT flag, status,set_for_release,release_date FROM ctf WHERE id = 2"
	}

	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var flag, status, setForRelease, releaseDate string

	if rows.Next() {
		err = rows.Scan(&flag, &status, &setForRelease, &releaseDate)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, err
	}

	settings := map[string]string{
		"flag":            flag,
		"status":          status,
		"set_for_release": setForRelease,
		"release_date":    releaseDate,
	}

	return settings, nil
}

func SaveSettings(status, setForRelease, flag, releaseDate, ctfType string) error {

	if ctfType == "jeopardy" {
		_, err := DB.Exec("UPDATE ctf SET status = $1, set_for_release = $2, flag = $3, release_date = $4 docker_image='none' WHERE id = 1", status, setForRelease, flag, releaseDate)
		if err != nil {
			return err
		}

	} else if ctfType == "attack-defense" {
		_, err := DB.Exec("UPDATE ctf SET status = $1, set_for_release = $2, flag = $3, release_date = $4, docker_image='none' WHERE id = 2", status, setForRelease, flag, releaseDate)
		if err != nil {
			return err
		}

	}

	return nil
}
