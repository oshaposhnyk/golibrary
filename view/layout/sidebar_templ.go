// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package layout

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"fmt"
	"math/rand"
)

func setLogo() string {
	randomNum := rand.Intn(4) + 1
	return fmt.Sprintf("/static/img/%d.png", randomNum)
}

func Sidebar() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<ul class=\"navbar-nav bg-gradient-info sidebar sidebar-dark accordion\" id=\"accordionSidebar\"><!-- Sidebar - Brand --><a class=\"sidebar-brand d-flex align-items-center justify-content-center\" href=\"/\"><div class=\"sidebar-brand-icon\"><img src=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(setLogo())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/layout/sidebar.templ`, Line: 19, Col: 40}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" height=\"80px\" alt=\"Golang Бібліотека\"></div><div class=\"sidebar-brand-text mx-3\">Golang<br>Library</div></a><!-- Divider --><hr class=\"sidebar-divider my-0\"><!-- Nav Item - Dashboard --><li class=\"nav-item\"><a class=\"nav-link\" href=\"/\"><i class=\"fas fa-fw fa-home\"></i> <span>Main</span></a></li><hr class=\"sidebar-divider d-none d-md-block\"><!-- Sidebar Toggler (Sidebar) --><div class=\"text-center d-none d-md-inline\"><button class=\"rounded-circle border-0\" id=\"sidebarToggle\"></button></div></ul><!-- End of Sidebar -->")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
