package callofduty

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	// The profileURL endpoint has five different parameters to retrieve
	// player profiles cross platform and for different games
	// version - API Version - Use v1 for MW, v2 for WWII, BO4, and CRM endpoints
	// game - Game - "mw" for Modern Warfare, "wwii" for WWII, "bo4" for Black Ops 4
	// platform - Platform associated with username - "uno" for Activision, "battle" for Battle.net, "steam" for WWII
	// username - URI encoded string of the platform-specific username (eg: MellowD#6992980 = MellowD%236992980)
	// mode - Game Mode - "wz" for Warzone or "mw" for Multiplayer
	profileURLTemplate = "https://my.callofduty.com/api/papi-client/stats/cod/%s/title/%s/platform/%s/gamer/%s/profile/type/%s"
)

// GetProfile retrieves a simplified profile
func (c *Client) GetProfile(game, platform, username, mode string) (ProfileData, error) {
	// Create default Profile
	profileData := ProfileData{}

	// Get the raw player profile response
	profileResponse, err := c.GetProfileRaw(game, platform, username, mode)
	if err != nil {
		return profileData, err
	}

	profileData = ProfileData(profileResponse.Data)

	return profileData, nil
}

// GetBrAllModeStats retrieves a player's profile and returns statistics
func (c *Client) GetBrAllModeStats(game, platform, username, mode string) (BrAllMode, error) {
	// Create an default empty BRStats type to return
	brAllModeStats := BrAllMode{}

	// Get the raw player profile response
	profileResponse, err := c.GetProfileRaw(game, platform, username, mode)
	if err != nil {
		return brAllModeStats, err
	}
	// Copy over data
	brAllModeStats = BrAllMode(profileResponse.Data.Lifetime.Mode.BrAll)

	return brAllModeStats, nil
}

// GetProfileRaw retrieves a player's profile based on
// the game, platform, username and mode
func (c *Client) GetProfileRaw(game, platform, username, mode string) (*Profile, error) {
	// Create default ProfileResponse
	profResp := &Profile{}

	// Set API version based on game value
	var apiVersion string
	if game == "mw" {
		apiVersion = "v1"
	} else {
		apiVersion = "v2"
	}

	// Create a new http.Client and construct the request
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	profileURL := fmt.Sprintf(profileURLTemplate, apiVersion, game, platform, username, mode)
	req, err := http.NewRequest("GET", profileURL, nil)
	if err != nil {
		return profResp, err
	}
	req.AddCookie(c.Authentication.ATKN)
	req.AddCookie(c.Authentication.SSO)
	req.AddCookie(c.Authentication.SSOExpiry)

	// Execute the request to get a response
	resp, err := client.Do(req)
	if err != nil {
		return profResp, err
	}
	defer resp.Body.Close()

	// Unmarshal the response body into a ProfileResponse
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return profResp, err
	}
	json.Unmarshal(body, profResp)
	return profResp, nil
}
