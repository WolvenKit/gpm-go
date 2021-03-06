/*
 Copyright (c) 2020 - 2021 the WolvenKit contributors.

 Licensed under the GNU Affero General Public License v3.0 (the "License").

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package game

import "go.uber.org/zap"

type Game struct {
    Directories       Directories
	Identifier        string
	DisplayName       string
	Versions          []string
	InstallStrategies []InstallStrategy
	ModRegistries     []ModRegistry
}

type Directories struct {
    GameRoot string
}

func InitGame(logger *zap.SugaredLogger) *Game {
    g := new(Game)
    return g
}
