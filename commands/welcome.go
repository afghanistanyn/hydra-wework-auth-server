package commands

import (
	"fmt"
	"github.com/mix-go/console"
	"github.com/mix-go/console/flag"
	"runtime"
	"strings"
)

const logo = `
.__              .___                                                  __                    __  .__                                              
|  |__ ___.__. __| _/___________    __  _  __ ______  _  _____________|  | __ _____   __ ___/  |_|  |__     ______ ______________  __ ___________ 
|  |  <   |  |/ __ |\_  __ \__  \   \ \/ \/ // __ \ \/ \/ /  _ \_  __ \  |/ / \__  \ |  |  \   __\  |  \   /  ___// __ \_  __ \  \/ // __ \_  __ \
|   Y  \___  / /_/ | |  | \// __ \_  \     /\  ___/\     (  <_> )  | \/    <   / __ \|  |  /|  | |   Y  \  \___ \\  ___/|  | \/\   /\  ___/|  | \/
|___|  / ____\____ | |__|  (____  /   \/\_/  \___  >\/\_/ \____/|__|  |__|_ \ (____  /____/ |__| |___|  / /____  >\___  >__|    \_/  \___  >__|   
     \/\/         \/            \/               \/                        \/      \/                 \/       \/     \/                 \/       

`

func welcome() {
	fmt.Println(strings.Replace(logo, "*", "`", -1))
	fmt.Println("")
	fmt.Println(fmt.Sprintf("Server      Name:      %s", "hydra-wework-auth-server"))
	fmt.Println(fmt.Sprintf("Listen      Addr:      %s", flag.Match("a", "addr").String(Addr)))
	fmt.Println(fmt.Sprintf("System      Name:      %s", runtime.GOOS))
	fmt.Println(fmt.Sprintf("Go          Version:   %s", runtime.Version()[2:]))
	fmt.Println(fmt.Sprintf("Framework   Version:   %s", console.Version))
}
