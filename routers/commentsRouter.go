package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["g8ink/controllers:AdminController"] = append(beego.GlobalControllerRouter["g8ink/controllers:AdminController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/`,
            AllowHTTPMethods: []string{"post","get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["g8ink/controllers:AdminController"] = append(beego.GlobalControllerRouter["g8ink/controllers:AdminController"],
        beego.ControllerComments{
            Method: "AddBan",
            Router: `/api/AddBan`,
            AllowHTTPMethods: []string{"post","get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["g8ink/controllers:AdminController"] = append(beego.GlobalControllerRouter["g8ink/controllers:AdminController"],
        beego.ControllerComments{
            Method: "AddWhiteList",
            Router: `/api/AddWhiteList`,
            AllowHTTPMethods: []string{"post","get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["g8ink/controllers:AdminController"] = append(beego.GlobalControllerRouter["g8ink/controllers:AdminController"],
        beego.ControllerComments{
            Method: "DeleteBan",
            Router: `/api/DeleteBan`,
            AllowHTTPMethods: []string{"post","get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["g8ink/controllers:AdminController"] = append(beego.GlobalControllerRouter["g8ink/controllers:AdminController"],
        beego.ControllerComments{
            Method: "DeleteLimitIp",
            Router: `/api/DeleteLimitIp`,
            AllowHTTPMethods: []string{"post","get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["g8ink/controllers:AdminController"] = append(beego.GlobalControllerRouter["g8ink/controllers:AdminController"],
        beego.ControllerComments{
            Method: "DeleteLink",
            Router: `/api/DeleteLink`,
            AllowHTTPMethods: []string{"post","get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["g8ink/controllers:AdminController"] = append(beego.GlobalControllerRouter["g8ink/controllers:AdminController"],
        beego.ControllerComments{
            Method: "DeleteWhiteList",
            Router: `/api/DeleteWhiteList`,
            AllowHTTPMethods: []string{"post","get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["g8ink/controllers:AdminController"] = append(beego.GlobalControllerRouter["g8ink/controllers:AdminController"],
        beego.ControllerComments{
            Method: "Ban",
            Router: `/ban`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["g8ink/controllers:AdminController"] = append(beego.GlobalControllerRouter["g8ink/controllers:AdminController"],
        beego.ControllerComments{
            Method: "Home",
            Router: `/home`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["g8ink/controllers:AdminController"] = append(beego.GlobalControllerRouter["g8ink/controllers:AdminController"],
        beego.ControllerComments{
            Method: "Limitips",
            Router: `/limitips`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["g8ink/controllers:AdminController"] = append(beego.GlobalControllerRouter["g8ink/controllers:AdminController"],
        beego.ControllerComments{
            Method: "Links",
            Router: `/links`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["g8ink/controllers:AdminController"] = append(beego.GlobalControllerRouter["g8ink/controllers:AdminController"],
        beego.ControllerComments{
            Method: "WhiteList",
            Router: `/whitelist`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
