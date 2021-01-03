/*
 Copyright (c) 2020 - 2021 the WolvenKit contributors.

 Licensed under the GNU Affero General Public License v3.0 (the "License").

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package test

import (
    "github.com/WolvenKit/gpm/internal/gpm/mod"
    "github.com/stretchr/testify/assert"
    "testing"
)

// User flow
/*
   User downloads mod to $downloadPath
   User unpacks mod archive into $downloadPath/tmp/
   GPM reads mods manifest in $downloadPath/tmp/$archiveName/manifest.toml
   User install mod from $downloadPath/tmp/$archiveName/**
   User enables mod
   User disables mod
   User uninstalls mod
*/

// Test downloads go to desired directory
func TestDownloadMod(t *testing.T) {
	tmp := createSandbox(true)
	logger := initLogging()

	m := mod.InitMod(logger)

	// Mock user CLI input
	i := new(mod.DownloadInput)
	i.Url = "https://cybermods.net/package/download/osulli/BraindanceProtocol/0.4.0/"
	i.Identifier = "braindance_protocol"
	i.FileType = ".zip"

	// Download Mod
	m.Download(logger, tmp, i)

	// Assert the mod archive was downloaded into the correct path
	assert.FileExists(t, m.Directories.ArchivePath)
}

// Tests mod Configuration is read correctly
func TestReadModConfiguration(t *testing.T) {
	logger := initLogging()

	m := mod.InitMod(logger)
	m.Directories.InstallDirectory = "mocks/example_cet_mod"
	m.ReadModConfiguration(logger)
	logger.Debug(m.Creator)

	assert.Equal(t, mod.Mod{
		Creator:           "WolvenKit",
		Identifier:        "braindance-protocol",
		Version:           "0.0.0",
		DisplayName:       "Braindance Protocol",
		Description:       "A collection of LUA scripts to modify your Cyberpunk 2077 experience",
		License:           "GNU v3",
		WebsiteURL:        "https://github.com/WolvenKit/BraindanceProtocol/",
		Dependencies:      []string{""},
		Tags:              []string{""},
		//InstallStrategies: []string{"CET"},
		ExtraData:         []string{""},
	}, m)
}

//// Tests mod Install follows install strategy
//func TestInstallMod(t *testing.T) {
//	tmp := createSandbox(true)
//	logger := initLogging()
//
//	m := mod.InitMod(logger, "mocks/example_cet_mod")
//	m.ReadModConfiguration()
//
//	mod := cmd.InstallMod(logger, "mocks/example_cet_mod.rar", tmp, m.Identifier)
//	assert.DirExists(t, mod)
//}

// Ensure scenario where mod manifest has missing keys is handled
func ModManifestMissingKeys() {
	// Checks CET not existing handled etc.
}

// Ensure scenario where mod archive cannot be found is handled
func UnarchiveMissingArchive() {
	// Checks CET not existing handled etc.
}

// Ensure scenario is handled if required directories do not exist
func InstallToInvalidDirectory() {
	// Checks CET not existing handled etc.
}

func InstallMod(t *testing.T) {
}

func UninstallMod(t *testing.T) {
}

func EnableMod(t *testing.T) {
}

func DisableMod(t *testing.T) {
}
