package adminlte

var TemplateList = map[string]string{"admin_panel":`{{define "admin_panel"}}
    <div class="navbar-custom-menu">
        <ul class="nav navbar-nav">

            {{.NavButtonsHTML}}

            <li title="{{lang "Fixed the sidebar"}}">
                <a href="javascript:void(0);" class="fixed-btn" data-click="false">
                    <i class="fa fa-thumb-tack"></i>
                </a>
            </li>

            <li title="{{lang "Enter fullscreen"}}">
                <a href="javascript:void(0);" class="fullpage-btn">
                    <i class="fa fa-arrows-alt"></i>
                </a>
            </li>
            <li title="{{lang "Exit fullscreen"}}">
                <a href="javascript:void(0);" class="exit-fullpage-btn" style="display: none;">
                    <i class="fa fa-compress"></i>
                </a>
            </li>
            <li title="{{lang "Refresh"}}">
                <a href="javascript:void(0);" class="container-refresh">
                    <i class="fa fa-refresh"></i>
                </a>
            </li>
            <!-- User Account: style can be found in dropdown.less -->
            <li class="dropdown user user-menu">
                <a href="#" class="dropdown-toggle" data-toggle="dropdown">
                    {{if eq .User.Avatar ""}}
                        <img src="{{.UrlPrefix}}/assets/dist/img/avatar04.png" class="user-image" alt="User Image">
                    {{else}}
                        <img src="{{.User.Avatar}}" class="user-image" alt="User Image">
                    {{end}}
                    <span class="hidden-xs">{{.User.Name}}</span>
                </a>
                <ul class="dropdown-menu">
                    <!-- User image -->
                    <li class="user-header">
                        {{if eq .User.Avatar ""}}
                            <img src="{{.UrlPrefix}}/assets/dist/img/avatar04.png" class="img-circle" alt="User Image">
                        {{else}}
                            <img src="{{.User.Avatar}}" class="img-circle" alt="User Image">
                        {{end}}
                        <p>
                            {{.User.Name}} -{{.User.LevelName}}
                            <small>{{.User.CreatedAt}}</small>
                        </p>
                    </li>
                    <li class="user-footer">
                        <div class="pull-left">
                            <a href="{{.UrlPrefix}}/info/normal_manager/edit?__goadmin_edit_pk={{.User.Id}}"
                               class="btn btn-default btn-flat">{{lang "setting"}}</a>
                        </div>
                        <div class="pull-right">
                            <a href="{{.UrlPrefix}}/logout" class="btn btn-default btn-flat">{{lang "sign out"}}</a>
                        </div>
                    </li>
                </ul>
            </li>
            <!-- <li class="hidden-xs">
                 <a href="javascript:;" data-toggle="control-sidebar"><i class="fa fa-gears"></i></a>
             </li> -->
        </ul>
    </div>
{{end}}`,"components/alert":`{{define "alert"}}
<div class="alert alert-{{.Theme}} alert-dismissible">
    <button type="button" class="close" data-dismiss="alert" aria-hidden="true">×</button>
    <h4>{{langHtml .Title}}</h4>
    {{langHtml .Content}}
</div>
{{end}}`,"components/box":`{{define "box"}}
<div class="box box-{{.Theme}}">
    {{if eq .HeadColor ""}}
        <div class="box-header {{.HeadBorder}}">
    {{else}}
        <div class="box-header {{.HeadBorder}}" style="background-color: {{.HeadColor}};">
    {{end}}
        {{langHtml .Header}}
    </div>
    {{if ne .SecondHeader ""}}
        {{if eq .SecondHeadColor ""}}
            <div class="box-header {{.SecondHeaderClass}} {{.SecondHeadBorder}}">
        {{else}}
            <div class="box-header {{.SecondHeaderClass}} {{.SecondHeadBorder}}" style="background-color: {{.SecondHeadColor}};">
        {{end}}
            {{langHtml .SecondHeader}}
        </div>
    {{end}}
    <div class="box-body" {{.Style}}>
        {{langHtml .Body}}
    </div>
    {{if ne .Footer ""}}
    <div class="box-footer clearfix">
        {{langHtml .Footer}}
    </div>
    {{end}}
</div>
{{end}}`,"components/button":`{{define "button"}}
    <div class="btn-group {{.Orientation}}" {{.Style}}>
        {{if eq .Href ""}}
            {{if ne .LoadingText ""}}
                <button type="{{.Type}}" class="btn {{.Size}} btn-{{.Theme}}" data-loading-text="{{.LoadingText}}">
            {{else}}
                <button type="{{.Type}}" class="btn {{.Size}} btn-{{.Theme}}">
            {{end}}
                {{langHtml .Content}}
            </button>
        {{else}}
            {{if ne .LoadingText ""}}
                <a href="{{.Href}}" type="{{.Type}}" class="btn {{.Size}} btn-{{.Theme}}" data-loading-text="{{.LoadingText}}">
            {{else}}
                <a href="{{.Href}}" type="{{.Type}}" class="btn {{.Size}} btn-{{.Theme}}">
            {{end}}
                {{langHtml .Content}}
            </a>
        {{end}}
    </div>
{{end}}`,"components/col":`{{define "col"}}
<div class="{{.Size}}">{{langHtml .Content}}</div>
{{end}}`,"components/form/color":`{{define "form_color"}}
    <div class="input-group colorpicker-element">
        <span class="input-group-addon"><i style="background-color: rgb(0, 0, 0);"></i></span>
        <input {{if .Must}}required="1"{{end}} style="width: 140px" type="text" id="{{.Field}}" name="{{.Field}}"
               value="" class="form-control {{.Field}}" placeholder="{{.Value}}">
    </div>
    <script>
        $('.{{.Field}}').parent().colorpicker([]);
    </script>
{{end}}`,"components/form/currency":`{{define "form_currency"}}
    <div class="input-group">
        <span class="input-group-addon">$</span>
        <input {{if .Must}}required="1"{{end}} style="width: 120px; text-align: right;" type="text" id="{{.Field}}"
               name="{{.Field}}"
               value="{{.Value}}" class="form-control {{.Field}}" placeholder="{{.Head}}">
    </div>
    <script>
        $(function () {
            $('.{{.Field}}').inputmask({
                "alias": "currency",
                "radixPoint": ".",
                "prefix": "",
                "removeMaskOnSubmit": true
            });
        });
    </script>
{{end}}`,"components/form/custom":`{{define "form_custom"}}
    <div class="input-group">
        {{.CustomContent}}
    </div>
    {{if .CustomJs}}
        <script>
            {{.CustomJs}}
        </script>
    {{end}}
    {{if .CustomCss}}
        <style>
            {{.CustomCss}}
        </style>
    {{end}}
{{end}}`,"components/form/datetime":`{{define "form_datetime"}}
    {{if not .Editable}}
        <div class="box box-solid box-default no-margin">
            <div class="box-body" style="min-height: 40px;">
                {{.Value}}
            </div>
            <input type="hidden" name="{{.Field}}" value='{{.Value}}'>
        </div>
    {{else}}
        <div class="input-group">
            {{if ne .Label ""}}
                <span class="input-group-addon">{{.Label}}</span>
            {{end}}
            <span class="input-group-addon"><i class="fa fa-calendar fa-fw"></i></span>
            <input {{if .Must}}required="1"{{end}} style="width: 160px" type="text" id="{{.Field}}"
                   name="{{.Field}}"
                   value="{{.Value}}"
                   class="form-control {{.Field}}" placeholder="{{.Placeholder}}">
        </div>
        <script>
            $(function () {
                $('.{{.Field}}').parent().datetimepicker({
                    "format": "YYYY-MM-DD HH:mm:ss",
                    "locale": "zh-CN",
                    "allowInputToggle": true
                });
            });
        </script>
    {{end}}
{{end}}`,"components/form/datetime_range":`{{define "form_datetime_range"}}
    <div class="input-group">
        {{if ne .Label ""}}
            <span class="input-group-addon">{{.Label}}</span>
        {{end}}
        <span class="input-group-addon"><i class="fa fa-calendar fa-fw"></i></span>
        <input type="text" id="{{.Field}}_start__goadmin" name="{{.Field}}_start__goadmin" value="{{.Value}}"
               class="form-control {{.Field}}_start__goadmin" placeholder="{{.Placeholder}}">
        <span class="input-group-addon" style="border-left: 0; border-right: 0;">-</span>
        <input type="text" id="{{.Field}}_end__goadmin" name="{{.Field}}_end__goadmin" value="{{.Value2}}"
               class="form-control {{.Field}}_end__goadmin" placeholder="{{.Placeholder}}">
    </div>
    <script>
        $(function () {
            $('.{{.Field}}_start__goadmin').datetimepicker({"format": "YYYY-MM-DD HH:mm:ss", "locale": "zh-CN"});
            $('.{{.Field}}_end__goadmin').datetimepicker({
                "format": "YYYY-MM-DD HH:mm:ss",
                "locale": "zh-CN",
                "useCurrent": false
            });
            $('.{{.Field}}_start__goadmin').on("dp.change", function (e) {
                $('.{{.Field}}_end__goadmin').data("DateTimePicker").minDate(e.date);
            });
            $('.{{.Field}}_end__goadmin').on("dp.change", function (e) {
                $('.{{.Field}}_start__goadmin').data("DateTimePicker").maxDate(e.date);
            });
        });
    </script>
{{end}}`,"components/form/default":`{{define "form_default"}}
    <div class="box box-solid box-default no-margin">
        <div class="box-body" style="min-height: 40px;">
            {{.Value}}
        </div>
    </div>
{{end}}`,"components/form/email":`{{define "form_email"}}
    <div class="input-group">
        <span class="input-group-addon"><i class="fa fa-envelope fa-fw"></i></span>
        <input {{if .Must}}required="1"{{end}} type="email" id="{{.Field}}" name="{{.Field}}" value='{{.Value}}'
               class="form-control json"
               placeholder="{{.Placeholder}}">
    </div>
{{end}}`,"components/form/file":`{{define "form_file"}}
    <input type="file" class="{{.Field}}" name="{{.Field}}" data-initial-preview="{{.Value2}}"
           data-initial-caption="{{.Value}}">
    <input type="hidden" value="0" name="{{.Field}}__delete_flag" class="{{.Field}}__delete_flag">
    <script>
        $("input.{{.Field}}").fileinput({{.OptionExt}});
        $(".preview-{{.Field}} .close.fileinput-remove").on("click", function (e) {
            $(".{{.Field}}__delete_flag").val("1")
        });
    </script>
{{end}}`,"components/form/help_block":`{{define "help_block"}}
    {{if ne . ""}}
        <span class="help-block">
            <i class="fa fa-info-circle"></i>&nbsp;{{.}}
        </span>
    {{end}}
{{end}}`,"components/form/iconpicker":`{{define "form_iconpicker"}}
    <div class="input-group">
        <span class="input-group-addon"><i class="fa"></i></span>
        {{if eq .Value ""}}
            <input style="width: 140px" type="text" name="{{.Field}}" value="fa-bars"
                   class="form-control {{.Field}}"
                   placeholder="{{lang "Input Icon"}}">
        {{else}}
            <input style="width: 140px" type="text" name="{{.Field}}" value="{{.Value}}"
                   class="form-control {{.Field}}"
                   placeholder="{{lang "Input Icon"}}">
        {{end}}
    </div>
    <script>
        $('.{{.Field}}').iconpicker({placement: 'bottomLeft'});
    </script>
{{end}}`,"components/form/ip":`{{define "form_ip"}}
    <div class="input-group">
        <span class="input-group-addon"><i class="fa fa-laptop fa-fw"></i></span>
        <input {{if .Must}}required="1"{{end}} style="width: 130px" type="text" id="{{.Field}}" name="{{.Field}}"
               value='{{.Value}}' class="form-control json"
               placeholder="{{.Placeholder}}">
    </div>
{{end}}`,"components/form/multi_file":`{{define "form_multi_file"}}
    <input type="file" class="{{.Field}}" name="{{.Field}}" multiple data-initial-caption="{{lang "Input"}} {{.Field}}">
    <input type="hidden" value="0" name="{{.Field}}__delete_flag" class="{{.Field}}__delete_flag">
    <script>
        mutilfileoptions = {{.OptionExt}};
        {{if ne .Value ""}}
        mutilfileoptions.initialPreview = {{js .Value}};
        {{end}}
        $("input.{{.Field}}").fileinput(mutilfileoptions);
        $(".preview-{{.Field}} .close.fileinput-remove").on("click", function (e) {
            $(".{{.Field}}__delete_flag").val("1")
        });
    </script>
{{end}}`,"components/form/number":`{{define "form_number"}}
    <div class="input-group">
        <input {{if .Must}}required="1"{{end}} style="width: 100px; text-align: center;" type="text" id="{{.Field}}"
               name="{{.Field}}"
               value="{{.Value}}" class="form-control {{.Field}}"
               placeholder="{{.Head}}">
    </div>
    <script>
        $(function () {
            $('.{{.Field}}:not(.initialized)')
                .addClass('initialized')
                .bootstrapNumber({
                    upClass: 'success',
                    downClass: 'primary',
                    center: true
                });
        })
    </script>
{{end}}`,"components/form/number_range":`{{define "form_number_range"}}
    <div class="input-group number-range">
        <input style="text-align: center;" type="text" id="{{.Field}}_start__goadmin"
               name="{{.Field}}_start__goadmin"
               value="{{.Value}}" class="form-control {{.Field}}_start__goadmin"
               placeholder="{{.Head}}">
        <span class="input-group-addon" style="border-left: 0; border-right: 0;">-</span>
        <input style="text-align: center;" type="text" id="{{.Field}}_end__goadmin" name="{{.Field}}_end__goadmin"
               value="{{.Value2}}" class="form-control {{.Field}}_end__goadmin"
               placeholder="{{.Head}}">
    </div>
    <script>
        $(function () {
            $('.{{.Field}}_start__goadmin:not(.initialized)')
                .addClass('initialized')
                .bootstrapNumber({
                    upClass: 'success',
                    downClass: 'primary',
                    center: true
                });
            $('.{{.Field}}_end__goadmin:not(.initialized)')
                .addClass('initialized')
                .bootstrapNumber({
                    upClass: 'success',
                    downClass: 'primary',
                    center: true
                });
        })
    </script>
    <style>
        .number-range .input-group {
            width: 100%;
        }
    </style>
{{end}}`,"components/form/password":`{{define "form_password"}}
    {{if .Editable}}
        <div class="input-group">
            <span class="input-group-addon"><i class="fa fa-eye-slash"></i></span>
            <input {{if .Must}}required="1"{{end}} type="password" id="{{.Field}}" name="{{.Field}}"
                   value="{{.Value}}"
                   class="form-control password" placeholder="{{.Placeholder}}">
        </div>
    {{else}}
        <div class="box box-solid box-default no-margin">
            <div class="box-body">********</div>
        </div>
    {{end}}
{{end}}`,"components/form/radio":`{{define "form_radio"}}
    {{$field := .Field}}
    {{range $key, $v := .Options }}
        <input type="radio" name="{{$field}}" value="{{$v.Value}}"
               class="minimal {{$field}}" {{$v.SelectedLabel}}
               style="position: absolute; opacity: 0;">&nbsp;{{$v.Text}}&nbsp;&nbsp;
    {{end}}
    <script>
        $(function () {
            $('.{{.Field}}').iCheck({radioClass: 'iradio_minimal-blue'});
        });
    </script>
{{end}}`,"components/form/richtext":`{{define "form_rich_text"}}
    <div id="{{.Field}}-editor">
    </div>
    <input type="hidden" id="{{.Field}}" name="{{.Field}}" value='{{.Value}}'
           placeholder="{{.Placeholder}}">
    <script type="text/javascript">
        {{$field := (js .Field)}}
        {{$field}}editor = new window.wangEditor('#{{.Field}}-editor');
        {{$field}}editor.customConfig.onchange = function (html) {
            $('#{{.Field}}').val(html)
        };
        {{.OptionExt}}
        {{$field}}editor.create();
        {{$field}}editor.txt.html('{{.Value}}');
        {{if not .Editable}}
        {{$field}}editor.$textElem.attr('contenteditable', false);
        {{end}}
        window.wangEditor.fullscreen.init('#{{.Field}}-editor');
    </script>
{{end}}`,"components/form/select":`{{define "form_select"}}
    <select class="form-control {{.Field}} select2-hidden-accessible" style="width: 100%;" name="{{.Field}}[]"
            multiple="" data-placeholder="{{.Placeholder}}" tabindex="-1" aria-hidden="true"
            {{if not .Editable}}disabled="disabled"{{end}}>
        <option></option>
        {{range $key, $v := .Options }}
            <option value='{{$v.Value}}' {{$v.SelectedLabel}}>{{$v.Text}}</option>
        {{end}}
    </select>
    <script>
        $(".{{.Field}}").select2({{.OptionExt}});
    </script>
{{end}}`,"components/form/selectbox":`{{define "form_selectbox"}}
    <select class="form-control {{.Field}}" style="width: 100%;" name="{{.Field}}[]" multiple="multiple"
            data-placeholder="Input {{.Head}}" {{if not .Editable}}disabled="disabled"{{end}}>
        {{range  $key, $v := .Options }}
            <option value='{{$v.Value}}' {{$v.SelectedLabel}}>{{$v.Text}}</option>
        {{end}}
    </select>
    <input type="hidden" name="{{.Field}}[]"/>
    <script>
        $(".{{.Field}}").bootstrapDualListbox({
            "infoText": "Showing all {0}",
            "infoTextEmpty": "Empty list",
            "infoTextFiltered": "{0} \/ {1}",
            "filterTextClear": "Show all",
            "filterPlaceHolder": "Filter"
        });
    </script>
{{end}}`,"components/form/singleselect":`{{define "form_select_single"}}
    <select class="form-control {{.Field}} select2-hidden-accessible" style="width: 100%;" name="{{.Field}}"
            data-multiple="false" data-placeholder="{{.Placeholder}}" tabindex="-1" aria-hidden="true"
            {{if not .Editable}}disabled="disabled"{{end}}>
        <option></option>
        {{range $key, $v := .Options }}
            <option value='{{$v.Value}}' {{$v.SelectedLabel}}>{{$v.Text}}</option>
        {{end}}
    </select>
    <script>
        $(".{{.Field}}").select2({{.OptionExt}});
    </script>
{{end}}`,"components/form/switch":`{{define "form_switch"}}
    <input id="__{{.Field}}" class="{{.Field}} ga_checkbox" {{(index .Options 0).SelectedLabel}} type="checkbox"
           name="__checkbox__{{.Field}}">
    {{$index := 0}}
    {{if eq (index .Options 0).SelectedLabel ""}}
        {{$index = 1}}
    {{end}}
    <input type="hidden" class="{{.Field}}" name="{{.Field}}" value="{{(index .Options $index).Value}}">
    <script>
        $(".{{.Field}}.ga_checkbox").bootstrapSwitch({
            size: "small",
            {{if eq (index .Options 0).Text ""}}
            onText: "ON",
            {{else}}
            onText: "{{(index .Options 0).Text}}",
            {{end}}
            {{if eq (index .Options 1).Text ""}}
            offText: "OFF",
            {{else}}
            offText: "{{(index .Options 1).Text}}",
            {{end}}
            onColor: 'primary',
            offColor: 'default',
            onSwitchChange: function (event, state) {
                $(event.target).closest('.bootstrap-switch').next().val(state ? '{{(index .Options 0).Value}}' : '{{(index .Options 1).Value}}').change();
            }
        })
    </script>
{{end}}`,"components/form/text":`{{define "form_text"}}
    {{if .Editable}}
        <div class="input-group">
            {{if eq .Label ""}}
                <span class="input-group-addon"><i class="fa fa-pencil fa-fw"></i></span>
            {{else if eq .Label "free"}}
                <div class="input-group-btn">
                    <input type="hidden" name="{{.Field}}__operator__" class="{{.Field}}-operation" value="3">
                    <button type="button" class="btn btn-default dropdown-toggle" data-toggle="dropdown"
                            style="min-width: 32px;" aria-expanded="false">
                        {{if eq .Value2 ""}}
                        <span class="{{.Field}}-label"> {{lang ">"}} </span>
                    {{else}}
                        <span class="{{.Field}}-label"> {{.Value2}} </span>
                    {{end}}&nbsp;&nbsp;
                        <span class="fa fa-caret-down"></span>
                    </button>
                    <ul class="dropdown-menu {{.Field}}_ul">
                        <li><a href="#" data-index="gr"> {{lang ">"}} </a></li>
                        <li><a href="#" data-index="le"> {{lang "<"}} </a></li>
                        <li><a href="#" data-index="gq"> {{lang ">="}} </a></li>
                        <li><a href="#" data-index="lq"> {{lang "<="}} </a></li>
                        <li><a href="#" data-index="eq"> {{lang "="}} </a></li>
                        <li><a href="#" data-index="ne"> {{lang "!="}} </a></li>
                    </ul>
                </div>
            {{else}}
                <span class="input-group-addon">{{.Label}}</span>
            {{end}}
            <input {{if .Must}}required="1"{{end}} type="text" id="{{.Field}}" name="{{.Field}}" value='{{.Value}}'
                   class="form-control json" placeholder="{{.Placeholder}}">
        </div>
    {{else}}
        <div class="box box-solid box-default no-margin">
            <div class="box-body">{{.Value}}</div>
        </div>
        <input type="hidden" name="{{.Field}}" value='{{.Value}}'>
    {{end}}
    {{if eq .Label "free"}}
        <script>
            $(function () {
                $(".{{.Field}}_ul li a").click(function () {
                    $(".{{.Field}}-label").text($(this).text());
                    $(".{{.Field}}-operation").val($(this).data('index'));
                });
            })
        </script>
    {{end}}
{{end}}`,"components/form/textarea":`{{define "form_textarea"}}
    <textarea {{if .Must}}required="1"{{end}} name="{{.Field}}" class="form-control" rows="5"
              placeholder="{{.Placeholder}}"
                      {{if not .Editable}}disabled="disabled"{{end}}>{{.Value}}</textarea>
{{end}}`,"components/form/url":`{{define "form_url"}}
    <div class="input-group">
        <span class="input-group-addon"><i class="fa fa-internet-explorer fa-fw"></i></span>
        <input {{if .Must}}required="1"{{end}} type="text" id="{{.Field}}" name="{{.Field}}" value='{{.Value}}'
               class="form-control json"
               placeholder="{{.Placeholder}}">
    </div>
{{end}}`,"components/form":`{{define "form"}}
    <script src={{link .CdnUrl .Prefix "/assets/dist/js/form.min.js"}}></script>
    {{.Header}}
    <form action="{{.Url}}" method="{{.Method}}" accept-charset="UTF-8" class="form-horizontal" pjax-container style="background-color: white;">
        <div class="box-body">

            {{if ne (len .TabHeaders) 0}}
                {{ template "form_layout_tab" . }}
            {{else if ne (len .ContentList) 0}}
                {{ template "form_layout_two_col" . }}
            {{else if .Layout.Flow}}
                {{ template "form_layout_flow" . }}
            {{else}}
                {{ template "form_layout_default" . }}
            {{end}}

        </div>

        {{if ne .OperationFooter ""}}
            <div class="box-footer">
                {{.OperationFooter}}
            </div>
        {{end}}

        {{range $key, $value := .HiddenFields}}
            <input type="hidden" name="{{$key}}" value='{{$value}}'>
        {{end}}
    </form>
    {{.Footer}}
{{end}}`,"components/form_components":`{{define "form_components"}}
    {{if eq .FormType.String "default"}}
        {{ template "form_default" .  }}
    {{else if eq .FormType.String "text"}}
        {{ template "form_text" .  }}
    {{else if eq .FormType.String "file"}}
        {{ template "form_file" .  }}
    {{else if eq .FormType.String "multi_file"}}
        {{ template "form_multi_file" .  }}
    {{else if eq .FormType.String "password"}}
        {{ template "form_password" .  }}
    {{else if eq .FormType.String "selectbox"}}
        {{ template "form_selectbox" .  }}
    {{else if eq .FormType.String "select"}}
        {{ template "form_select" .  }}
    {{else if eq .FormType.String "select_single"}}
        {{ template "form_select_single" .  }}
    {{else if eq .FormType.String "textarea"}}
        {{ template "form_textarea" .  }}
    {{else if eq .FormType.String "iconpicker"}}
        {{ template "form_iconpicker" .  }}
    {{else if eq .FormType.String "richtext"}}
        {{ template "form_rich_text" .  }}
    {{else if eq .FormType.String "datetime"}}
        {{ template "form_datetime" .  }}
    {{else if eq .FormType.String "datetime_range"}}
        {{ template "form_datetime_range" .  }}
    {{else if eq .FormType.String "radio"}}
        {{ template "form_radio" .  }}
    {{else if eq .FormType.String "email"}}
        {{ template "form_email" .  }}
    {{else if eq .FormType.String "url"}}
        {{ template "form_url" .  }}
    {{else if eq .FormType.String "ip"}}
        {{ template "form_ip" .  }}
    {{else if eq .FormType.String "color"}}
        {{ template "form_color" .  }}
    {{else if eq .FormType.String "currency"}}
        {{ template "form_currency" .  }}
    {{else if eq .FormType.String "number"}}
        {{ template "form_number" .  }}
    {{else if eq .FormType.String "number_range"}}
        {{ template "form_number_range" .  }}
    {{else if eq .FormType.String "custom"}}
        {{ template "form_custom" .  }}
    {{else if eq .FormType.String "switch"}}
        {{ template "form_switch" .  }}
    {{end}}
    {{if ne .HelpMsg ""}}
        <span class="help-block">
            <i class="fa fa-info-circle"></i>&nbsp;{{.HelpMsg}}
        </span>
    {{end}}
{{end}}`,"components/form_components_layout":`{{define "form_components_layout"}}

    {{if ne (len .ContentList) 0}}

        {{range $key, $content := .ContentList}}
            <div class="col-md-{{divide 12 (len $.ContentList)}}">
                <div class="box-body">
                    <div class="fields-group">
                        {{range $key, $data := $content}}
                            {{if $data.Hide}}
                                <input type="hidden" name="{{$data.Field}}" value='{{$data.Value}}'>
                            {{else}}
                                <div class="form-group" {{if ne $data.Width 0}}style="width: {{$data.Width}}px;"{{end}}>
                                    {{if ne $data.Head ""}}
                                        <label for="{{$data.Field}}"
                                               class="col-sm-{{$.HeadWidth}} {{if $data.Must}}asterisk{{end}} control-label">{{$data.Head}}</label>
                                    {{end}}
                                    <div class="col-sm-{{$.InputWidth}}">
                                        {{template "form_components" $data}}
                                    </div>
                                </div>
                            {{end}}
                        {{end}}
                    </div>
                </div>
            </div>
        {{end}}

    {{else if ne (len .TabHeaders) 0}}

        {{range $key, $content := .TabContents}}
            <div class="tab-pane {{if eq $key 0}}active{{end}}" id="tab-form-{{$key}}">
                {{range $key, $data := $content}}
                    {{if $data.Hide}}
                        <input type="hidden" name="{{$data.Field}}" value='{{$data.Value}}'>
                    {{else}}
                        <div class="form-group" {{if ne $data.Width 0}}style="width: {{$data.Width}}px;"{{end}}>
                            {{if ne $data.Head ""}}
                                <label for="{{$data.Field}}"
                                       class="col-sm-{{$.HeadWidth}} {{if $data.Must}}asterisk{{end}} control-label">{{$data.Head}}</label>
                            {{end}}
                            <div class="col-sm-{{$.InputWidth}}">
                                {{template "form_components" $data}}
                            </div>
                        </div>
                    {{end}}
                {{end}}
                {{range $key, $d := $content}}
                    {{if eq $d.Field $.PrimaryKey}}
                        <input type="hidden" name="{{$.PrimaryKey}}" value='{{$d.Value}}'>
                    {{end}}
                {{end}}
            </div>
        {{end}}

    {{else if .Layout.Flow}}

        {{range $key, $data := .Content}}
            {{if $data.Hide}}
                <input type="hidden" name="{{$data.Field}}" value='{{$data.Value}}'>
            {{else}}
                <div class="form-group" style="float: left;{{if ne $data.Width 0}}width: {{$data.Width}}px;{{end}}">
                    {{if ne $data.Head ""}}
                        <label for="{{$data.Field}}"
                               class="col-sm-{{$.HeadWidth}} {{if $data.Must}}asterisk{{end}} control-label">{{$data.Head}}</label>
                    {{end}}
                    <div class="col-sm-{{$.InputWidth}}">
                        {{template "form_components" $data}}
                    </div>
                </div>
            {{end}}
        {{end}}

    {{else}}

        {{range $key, $data := .Content}}
            {{if $data.Hide}}
                <input type="hidden" name="{{$data.Field}}" value='{{$data.Value}}'>
            {{else}}
                <div class="form-group" {{if ne $data.Width 0}}style="width: {{$data.Width}}px;"{{end}}>
                    {{if ne $data.Head ""}}
                        <label for="{{$data.Field}}"
                               class="col-sm-{{$.HeadWidth}} {{if $data.Must}}asterisk{{end}} control-label">{{$data.Head}}</label>
                    {{end}}
                    <div class="col-sm-{{$.InputWidth}}">
                        {{template "form_components" $data}}
                    </div>
                </div>
            {{end}}
        {{end}}

    {{end}}
{{end}}`,"components/form_layout_default":`{{define "form_layout_default"}}

    <div class="box-body">
        <div class="fields-group">
            {{ template "form_components_layout" . }}
        </div>

        {{range $key, $data := .Content}}
            {{if eq $data.Field $.PrimaryKey}}
                <input type="hidden" name="{{$.PrimaryKey}}" value='{{$data.Value}}'>
            {{end}}
        {{end}}
    </div>

{{end}}`,"components/form_layout_flow":`{{define "form_layout_flow"}}

    <div class="box-body">
        <div class="fields-group">
            {{ template "form_components_layout" . }}
        </div>

        {{range $key, $data := .Content}}
            {{if eq $data.Field $.PrimaryKey}}
                <input type="hidden" name="{{$.PrimaryKey}}" value='{{$data.Value}}'>
            {{end}}
        {{end}}
    </div>

{{end}}`,"components/form_layout_tab":`{{define "form_layout_tab"}}

    <div class="box-body">
        <div class="nav-tabs-custom">
            <ul class="nav nav-tabs">
                {{range $key, $data := .TabHeaders}}
                    {{if eq $key 0}}
                        <li class="active">
                            <a href="#tab-form-{{$key}}" data-toggle="tab" aria-expanded="true">
                                {{$data}} <i class="fa fa-exclamation-circle text-red hide"></i>
                            </a>
                        </li>
                    {{else}}
                        <li class="">
                            <a href="#tab-form-{{$key}}" data-toggle="tab" aria-expanded="true">
                                {{$data}} <i class="fa fa-exclamation-circle text-red hide"></i>
                            </a>
                        </li>
                    {{end}}
                {{end}}
            </ul>
            <div class="tab-content fields-group">
                {{ template "form_components_layout" .}}
            </div>
        </div>
    </div>

{{end}}`,"components/form_layout_two_col":`{{define "form_layout_two_col"}}
    <div class="row">
        {{ template "form_components_layout" .}}
    </div>

    {{range $key, $data := .Content}}
        {{if eq $data.Field $.PrimaryKey}}
            <input type="hidden" name="{{$.PrimaryKey}}" value='{{$data.Value}}'>
        {{end}}
    {{end}}
{{end}}`,"components/image":`{{define "image"}}
    {{if .HasModal}}
        <img src="{{.Src}}" width="{{.Width}}" height="{{.Height}}" data-toggle="modal" data-target="#img_{{.Uuid}}" style="cursor: zoom-in;">
        <div id="img_{{.Uuid}}" class="modal fade {{.Uuid}}" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true">
            <div class="modal-dialog {{.Uuid}}">
                <div class="modal-content {{.Uuid}}">
                    <div class="modal-body">
                        <img src="{{.Src}}" class="img-responsive">
                    </div>
                </div>
            </div>
        </div>
        <script>
            function centerModal() {
                $(this).css('display', 'block');
                var $dialog = $(this).find(".modal-dialog.{{.Uuid}}");
                var offset = ($(window).height() - $dialog.height()) / 2;
                $dialog.css("margin-top", offset);
            }

            $('.modal.{{.Uuid}}').on('show.bs.modal', centerModal);
            $(window).on("resize", function () {
                $('.modal:visible').each(centerModal);
            });
        </script>
    {{else}}
        <img src="{{.Src}}" width="{{.Width}}" height="{{.Height}}">
    {{end}}
{{end}}`,"components/label":`{{define "label"}}
<span class="label label-{{.Type}}" style="background-color: {{.Color}};">{{langHtml .Content}}</span>
{{end}}`,"components/link":`{{define "link"}}
    <a {{if .NewTab}}class="new-tab-link"{{end}} data-title="{{.Title}}" href="{{.URL}}">{{.Content}}</a>
{{end}}`,"components/paginator":`{{define "paginator"}}
    <div style="float: left;margin-top: 21px;">{{lang "showing"}} <b>{{.CurPageStartIndex}}</b> {{lang "to"}}
        <b>{{.CurPageEndIndex}}</b> {{lang "of"}} <b>{{.Total}}</b> {{lang "entries"}} &nbsp;&nbsp;&nbsp;{{.ExtraInfo}}
    </div>
    <ul class="pagination pagination-sm no-margin pull-right">
        <!-- Previous Page Link -->
        <li class="page-item {{.PreviousClass}}">
            {{if eq .PreviousClass "disabled"}}
                <span class="page-link">«</span>
            {{else}}
                <a class="page-link" href='{{.PreviousUrl}}' rel="next">«</a>
            {{end}}
        </li>

        <!-- Array Of Links -->
        {{range $key, $page := .Pages}}
            {{if eq (index $page "isSplit") "0"}}
                {{if eq (index $page "active") "active"}}
                    <li class="page-item active"><span class="page-link">{{index $page "page"}}</span></li>
                {{else}}
                    <li class="page-item"><a class="page-link" href='{{index $page "url"}}'>{{index $page "page"}}</a>
                    </li>
                {{end}}
            {{else}}
                <li class="page-item disabled"><span class="page-link">...</span></li>
            {{end}}
        {{end}}


        <!-- Next Page Link -->
        <li class='page-item {{.NextClass}}'>
            {{if eq .NextClass "disabled"}}
                <span class="page-link">»</span>
            {{else}}
                <a class="page-link" href='{{.NextUrl}}' rel="next">»</a>
            {{end}}
        </li>
    </ul>

    <label class="control-label pull-right" style="margin-right: 10px; font-weight: 100;">

        <small>{{lang "show"}}</small>&nbsp;
        {{$option := .Option}}
        {{$url := .Url}}
        <select class="input-sm grid-per-pager" name="per-page">
            {{range $key, $pageSize := .PageSizeList}}
                <option value="{{$url}}&__pageSize={{$pageSize}}" {{index $option $pageSize}}>
                    {{$pageSize}}
                </option>
            {{end}}
        </select>
        <small>{{lang "entries"}}</small>
    </label>

    <script>
        let gridPerPaper = $('.grid-per-pager');
        gridPerPaper.on('change', function () {
            $.pjax({url: this.value, container: '#pjax-container'});
        });
    </script>
{{end}}`,"components/popup":`{{define "popup"}}
<div class="modal fade" id="{{.ID}}" tabindex="-1" role="dialog" aria-labelledby="{{.ID}}" aria-hidden="true">
    <div class="modal-dialog modal-{{.Size}}" role="document">
        <div class="modal-content">
            <div class="modal-header">
                <h5 class="modal-title" id="{{.ID}}Title">{{langHtml .Title}}</h5>
                <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                    <span aria-hidden="true">&times;</span>
                </button>
            </div>
            <div class="modal-body">
                {{langHtml .Body}}
            </div>
            <div class="modal-footer">
                <button type="button" class="btn btn-secondary" data-dismiss="modal">{{lang "Close"}}</button>
                {{if .Footer}}
                <button type="button" class="btn btn-primary">{{langHtml .Footer}}</button>
                {{end}}
            </div>
        </div>
    </div>
</div>
{{end}}`,"components/row":`{{define "row"}}
<div class="row">{{langHtml .Content}}</div>
{{end}}`,"components/table/box-header":`{{define "box-header"}}
    <div class="pull-right">

        {{if ne .IsHideRowSelector true}}
            <div class="dropdown pull-right column-selector" style="margin-right: 10px">
                <button type="button" class="btn btn-sm btn-instagram dropdown-toggle" data-toggle="dropdown">
                    <i class="fa fa-table"></i>
                    &nbsp;
                    <span class="caret"></span>
                </button>
                <ul class="dropdown-menu" role="menu" style="padding: 10px;max-height: 400px;overflow: scroll;">
                    <li>
                        <ul style="padding: 0;">
                            {{range $key, $head := .Thead}}
                                <li class="checkbox icheck" style="margin: 0;">
                                    <label style="width: 100%;padding: 3px;">
                                        <input type="checkbox" class="column-select-item" data-id="{{$head.Field}}"
                                               style="position: absolute; opacity: 0;">&nbsp;&nbsp;&nbsp;{{$head.Head}}
                                    </label>
                                </li>
                            {{end}}
                        </ul>
                    </li>
                    <li class="divider">
                    </li>
                    <li class="text-right">
                        <button class="btn btn-sm btn-default column-select-all">{{lang "all"}}</button>&nbsp;&nbsp;
                        <button class="btn btn-sm btn-primary column-select-submit">{{lang "submit"}}</button>
                    </li>
                </ul>
            </div>
        {{end}}

        {{if .HasFilter}}

            <div class="btn-group pull-right" style="margin-right: 10px">
                <a href="javascript:;" class="btn btn-sm btn-primary" id="filter-btn"><i
                            class="fa fa-filter"></i>&nbsp;&nbsp;{{lang "filter"}}</a>
            </div>

            <script>
                $("#filter-btn").click(function () {
                    $('.filter-area').toggle();
                });
            </script>

        {{end}}

        <div class="btn-group pull-right" style="margin-right: 10px">
            {{if .NewUrl}}
                <a href="{{.NewUrl}}" class="btn btn-sm btn-success">
                    <i class="fa fa-plus"></i>&nbsp;&nbsp;{{lang "New"}}
                </a>
            {{end}}
            {{if .ExportUrl}}
                <div class="btn-group">
                    <a class="btn btn-sm btn-default">{{lang "Export"}}</a>
                    <button type="button" class="btn btn-sm btn-default dropdown-toggle" data-toggle="dropdown">
                        <span class="caret"></span>
                        <span class="sr-only">{{lang "Toggle Dropdown"}}</span>
                    </button>
                    <ul class="dropdown-menu" role="menu">
                        <li><a href="#" id="export-btn-0">{{lang "Current Page"}}</a></li>
                        {{if .ExportUrl}}
                            <li><a href="#" id="export-btn-1">{{lang "All"}}</a></li>
                        {{end}}
                    </ul>
                </div>
            {{end}}
        </div>
        {{renderRowDataHTML "" .Buttons}}
    </div>
    <span>
        {{if or .DeleteUrl .ExportUrl}}
            <div class="btn-group">
                <a class="btn btn-sm btn-default">{{lang "Action"}}</a>
                <button type="button" class="btn btn-sm btn-default dropdown-toggle" data-toggle="dropdown">
                <span class="caret"></span>
                <span class="sr-only">{{lang "Toggle Dropdown"}}</span>
                </button>
                <ul class="dropdown-menu" role="menu">
                    {{if .DeleteUrl}}
                        <li><a href="#" class="grid-batch-0">{{lang "Delete"}}</a></li>
                    {{end}}
                    {{if .ExportUrl}}
                        <li><a href="#" class="grid-batch-1">{{lang "Export"}}</a></li>
                    {{end}}
                </ul>
            </div>
        {{end}}
        <a class="btn btn-sm btn-primary grid-refresh">
            <i class="fa fa-refresh"></i> {{lang "Refresh"}}
        </a>
    </span>
    <script>
        let toastMsg = '{{lang "Refresh succeeded"}} !';
        $('.grid-refresh').on('click', function () {
            $.pjax.reload('#pjax-container');
            toastr.success(toastMsg);
        });

        {{if .ExportUrl}}

        $("#export-btn-0").click(function () {
            ExportData("false")
        });
        $("#export-btn-1").click(function () {
            ExportData("true")
        });

        function ExportData(isAll) {
            let form = $("<form>");
            form.attr("style", "display:none");
            form.attr("target", "");
            form.attr("method", "post");
            form.attr("action",{{.ExportUrl}});
            let input1 = $("<input>");
            input1.attr("type", "hidden");
            input1.attr("name", "time");
            input1.attr("value", (new Date()).getTime());
            let input2 = $("<input>");
            input2.attr("type", "hidden");
            input2.attr("name", "is_all");
            input2.attr("value", isAll);
            $("body").append(form);
            form.append(input1);
            form.append(input2);
            form.submit();
            form.remove()
        }

        {{end}}
    </script>
{{end}}`,"components/table":`{{define "table"}}
    <table class="table table-hover" style="min-width: {{.MinWidth}}px;table-layout: {{.Layout}};">
        {{if eq .Type "table"}}
            <thead>
            <tr>
                {{range $key, $head := .Thead}}
                    <th>
                        {{$head.Head}}
                    </th>
                {{end}}
            </tr>
            </thead>
        {{end}}
        <tbody>
        {{if eq .Type "data-table"}}
            <tr>
                {{if eq .IsTab false}}
                    <th style="text-align: center;">
                        <input type="checkbox" class="grid-select-all" style="position: absolute; opacity: 0;">
                    </th>
                {{end}}
                {{range $key, $head := .Thead}}
                    {{if eq $head.Hide false}}
                        {{if eq $head.Width 0}}
                            <th>
                        {{else}}
                            <th style="width: {{$head.Width}}px">
                        {{end}}
                        {{$head.Head}}
                        {{if $head.Sortable}}
                            <a class="fa fa-fw fa-sort" id="sort-{{$head.Field}}"
                               href="?__sort={{$head.Field}}&__sort_type=desc"></a>
                        {{end}}
                        </th>
                    {{end}}
                {{end}}
                {{if eq .NoAction false}}
                    <th style="text-align: center;">{{lang "operation"}}</th>
                {{end}}
            </tr>
        {{end}}


        {{$NoAction := .NoAction}}
        {{$Action := .Action}}
        {{$Thead := .Thead}}
        {{$Type := .Type}}
        {{$EditUrl := .EditUrl}}
        {{$UpdateUrl := .UpdateUrl}}
        {{$IsTab := .IsTab}}
        {{$DeleteUrl := .DeleteUrl}}
        {{$DetailUrl := .DetailUrl}}
        {{$PrimaryKey := .PrimaryKey}}
        {{range $key1, $info := .InfoList}}
            <tr>
                {{if eq $Type "data-table"}}
                    {{if eq $IsTab false}}
                        <td style="text-align: center;">
                            <input type="checkbox" class="grid-row-checkbox"
                                   data-id="{{(index $info $PrimaryKey).Content}}"
                                   style="position: absolute; opacity: 0;">
                        </td>
                    {{end}}
                {{end}}
                {{if eq $Type "data-table"}}
                    {{range $key2, $head2 := $Thead}}
                        {{if eq $head2.Hide false}}
                            {{if $head2.Editable}}
                                <td>
                                    {{if eq $head2.EditType "switch"}}
                                        <input class="info_edit_switch ga_checkbox"
                                               data-off-text="{{(index $head2.EditOption 1).Text}}"
                                               data-on-text="{{js (index $head2.EditOption 0).Text}}"
                                               data-size="{{index (index $head2.EditOption 0).Extra "size"}}"
                                               data-on-color="{{js (index (index $head2.EditOption 0).Extra "onColor")}}"
                                               data-off-color="{{index (index $head2.EditOption 0).Extra "offColor"}}"
                                               data-field="{{$head2.Field}}"
                                               data-updateurl="{{$UpdateUrl}}"
                                               data-pk="{{(index $info $PrimaryKey).Content}}"
                                               type="checkbox" name="__checkbox__edit_info"
                                                {{if eq (index $head2.EditOption 0).Value (index $info $head2.Field).Value}}
                                                    checked
                                                {{end}}
                                        >
                                        <input type="hidden" value="{{(index $head2.EditOption 0).Value}}">
                                        <input type="hidden" value="{{(index $head2.EditOption 1).Value}}">
                                    {{else}}
                                        <a href="#" class="editable-td-{{$head2.EditType}}"
                                           data-pk="{{(index $info $PrimaryKey).Content}}"
                                           data-source='{{$head2.EditOption.Marshal}}'
                                           data-url="{{$UpdateUrl}}"
                                           data-value="{{(index $info $head2.Field).Value}}"
                                           data-name="{{$head2.Field}}"
                                           data-title="Enter {{$head2.Head}}">{{(index $info $head2.Field).Content}}</a>
                                    {{end}}
                                </td>
                            {{else}}
                                <td>{{(index $info $head2.Field).Content}}</td>
                            {{end}}
                        {{end}}
                    {{end}}
                    {{if eq $NoAction false}}
                        <td style="text-align: center;">
                            {{if eq $Action ""}}
                                {{if $EditUrl}}
                                    <a href='{{$EditUrl}}&__goadmin_edit_pk={{(index $info $PrimaryKey).Content}}'><i
                                                class="fa fa-edit"></i></a>
                                {{end}}
                                {{if $DeleteUrl}}
                                    <a href="javascript:void(0);" data-id='{{(index $info $PrimaryKey).Content}}'
                                       class="grid-row-delete"><i class="fa fa-trash"></i></a>
                                {{end}}
                                {{if $DetailUrl}}
                                    <a href="{{$DetailUrl}}&__goadmin_detail_pk={{(index $info $PrimaryKey).Content}}"
                                       class="grid-row-view">
                                        <i class="fa fa-eye"></i>
                                    </a>
                                {{end}}
                            {{else}}
                                {{renderRowDataHTML (index $info $PrimaryKey).Content $Action}}
                            {{end}}
                        </td>
                    {{end}}
                {{else}}
                    {{range $key2, $head2 := $Thead}}
                        <td>{{(index $info $head2.Head).Content}}</td>
                    {{end}}
                {{end}}
            </tr>
        {{end}}
        </tbody>
    </table>
    {{if eq $Type "data-table"}}
        <script>
            window.selectedRows = function () {
                let selected = [];
                $('.grid-row-checkbox:checked').each(function () {
                    selected.push($(this).data('id'));
                });
                return selected;
            };

            const selectedAllFieldsRows = function () {
                let selected = [];
                $('.column-select-item:checked').each(function () {
                    selected.push($(this).data('id'));
                });
                return selected;
            };

            const pjaxContainer = "#pjax-container";
            const noAnimation = "__go_admin_no_animation_";

            function iCheck(el) {
                el.iCheck({checkboxClass: 'icheckbox_minimal-blue'}).on('ifChanged', function () {
                    if (this.checked) {
                        $(this).closest('tr').css('background-color', "#ffffd5");
                    } else {
                        $(this).closest('tr').css('background-color', '');
                    }
                });
            }

            $(function () {

                $('.grid-select-all').iCheck({checkboxClass: 'icheckbox_minimal-blue'}).on('ifChanged', function (event) {
                    console.log("event", event);
                    if (this.checked) {
                        $('.grid-row-checkbox').iCheck('check');
                    } else {
                        $('.grid-row-checkbox').iCheck('uncheck');
                    }
                });
                let items = $('.column-select-item');
                iCheck(items);
                iCheck($('.grid-row-checkbox'));
                let columns = getQueryVariable("__columns");
                if (columns === -1) {
                    items.iCheck('check');
                } else {
                    let columnsArr = columns.split(",");
                    for (let i = 0; i < columnsArr.length; i++) {
                        for (let j = 0; j < items.length; j++) {
                            if (columnsArr[i] === $(items[j]).attr("data-id")) {
                                $(items[j]).iCheck('check');
                            }
                        }
                    }
                }

                {{if .HasFilter}}{{if .IsHideFilterArea}}
                $('.filter-area').hide();
                {{end}}{{end}}

                // Fix PopUp error of table row action

                let lastTd = $("table tr:last td:last div");
                if (lastTd.hasClass("dropdown")) {
                    let popUpHeight = $("table tr:last td:last div ul").height();

                    let trs = $("table tr");
                    let totalHeight = 0;
                    for (let i = 1; i < trs.length - 1; i++) {
                        totalHeight += $(trs[i]).height();
                    }
                    if (popUpHeight > totalHeight) {
                        let h = popUpHeight + 16;
                        $("table tbody").append("<tr style='height:" + h + "px;'></tr>");
                    }

                    trs = $("table tr");
                    for (let i = trs.length - 1; i > 1; i--) {
                        let td = $(trs[i]).find("td:last-child div");
                        let combineHeight = $(trs[i]).height() / 2 - 20;
                        for (let j = i + 1; j < trs.length; j++) {
                            combineHeight += $(trs[j]).height();
                        }
                        if (combineHeight < popUpHeight) {
                            td.removeClass("dropdown");
                            td.addClass("dropup");
                        }
                    }
                }

                // Initialize sort parameters

                let sort = getQueryVariable("__sort");
                let sort_type = getQueryVariable("__sort_type");

                if (sort !== -1 && sort_type !== -1) {
                    let sortFa = $('#sort-' + sort);
                    if (sort_type === 'asc') {
                        sortFa.attr('href', '?__sort=' + sort + "&__sort_type=desc")
                    } else {
                        sortFa.attr('href', '?__sort=' + sort + "&__sort_type=asc")
                    }
                    sortFa.removeClass('fa-sort');
                    sortFa.addClass('fa-sort-amount-' + sort_type);
                }
            });

            // ============================
            // .IsHideRowSelector
            // ============================

            {{if ne .IsHideRowSelector true}}

            $('.column-select-all').on('click', function () {
                if ($(this).data('check') === '') {
                    $('.column-select-item').iCheck('check');
                    $(this).data('check', 'true')
                } else {
                    $('.column-select-item').iCheck('uncheck');
                    $(this).data('check', '')
                }
            });

            $('.column-select-submit').on('click', function () {

                let param = new Map();
                param.set('__columns', selectedAllFieldsRows().join(','));
                param.set(noAnimation, 'true');

                $.pjax({
                    url: addParameterToURL(param),
                    container: pjaxContainer
                });

                toastr.success('{{lang "reload succeeded"}} !');
            });

            {{end}}

            // ============================
            // end
            // ============================

            // ============================
            // .ExportUrl
            // ============================

            {{if .ExportUrl}}

            $('.grid-batch-1').on('click', function () {
                let rows = selectedRows();
                if (rows.length > 0) {
                    ExportAll(rows.join())
                }
            });

            function ExportAll(id) {
                let form = $("<form>");
                form.attr("style", "display:none");
                form.attr("target", "");
                form.attr("method", "post");
                form.attr("action",{{.ExportUrl}});
                let input1 = $("<input>");
                input1.attr("type", "hidden");
                input1.attr("name",{{.PrimaryKey}});
                input1.attr("value", id);
                $("body").append(form);
                form.append(input1);
                form.submit();
                form.remove()
            }

            {{end}}

            // ============================
            // end
            // ============================

            // ============================
            // .DeleteUrl
            // ============================

            {{if .DeleteUrl}}

            $('.grid-row-delete').unbind('click').click(function () {
                DeletePost($(this).data('id'))
            });

            $('.grid-batch-0').on('click', function () {
                let rows = selectedRows();
                if (rows.length > 0) {
                    DeletePost(rows.join())
                }
            });

            function DeletePost(id) {
                swal({
                        title: {{lang "are you sure to delete"}},
                        type: "warning",
                        showCancelButton: true,
                        confirmButtonColor: "#DD6B55",
                        confirmButtonText: {{lang "yes"}},
                        closeOnConfirm: false,
                        cancelButtonText: {{lang "cancel"}},
                    },
                    function () {
                        $.ajax({
                            method: 'post',
                            url: {{.DeleteUrl}},
                            data: {
                                id: id
                            },
                            success: function (data) {
                                let param = new Map();
                                param.set(noAnimation, "true");
                                $.pjax({
                                    url: addParameterToURL(param),
                                    container: pjaxContainer
                                });
                                if (typeof (data) === "string") {
                                    data = JSON.parse(data);
                                }
                                if (data.code === 200) {
                                    $('#_TOKEN').val(data.data);
                                    let lastTd = $("table tr:last td:last div");
                                    if (lastTd.hasClass("dropdown")) {
                                        let popUpHeight = $("table tr:last td:last div ul").height();

                                        let trs = $("table tr");
                                        let totalHeight = 0;
                                        for (let i = 1; i < trs.length - 1; i++) {
                                            totalHeight += $(trs[i]).height();
                                        }
                                        if (popUpHeight > totalHeight) {
                                            let h = popUpHeight + 16;
                                            $("table tbody").append("<tr style='height:" + h + "px;'></tr>");
                                        }
                                    }
                                    swal(data.msg, '', 'success');
                                } else {
                                    swal(data.msg, '', 'error');
                                }
                            }
                        });
                    });
            }

            {{end}}

            // ============================
            // end
            // ============================

            // ============================
            // Helper functions
            // ============================

            function getQueryVariable(variable) {
                let query = window.location.search.substring(1);
                let vars = query.split("&");
                for (let i = 0; i < vars.length; i++) {
                    let pair = vars[i].split("=");
                    if (pair[0] === variable) {
                        return pair[1];
                    }
                }
                return -1;
            }

            function addParameterToURL(params) {
                let newUrl = location.href.replace("#", "");

                for (let [field, value] of params) {
                    if (getQueryVariable(field) !== -1) {
                        newUrl = replaceParamVal(newUrl, field, value);
                    } else {
                        if (newUrl.indexOf("?") > 0) {
                            newUrl = newUrl + "&" + field + "=" + value;
                        } else {
                            newUrl = newUrl + "?" + field + "=" + value;
                        }
                    }
                }

                return newUrl
            }

            function replaceParamVal(oUrl, paramName, replaceWith) {
                let re = eval('/(' + paramName + '=)([^&]*)/gi');
                return oUrl.replace(re, paramName + '=' + replaceWith);
            }

            $(function () {

                $('.editable-td-select').editable({
                    "type": "select",
                    "emptytext": "<i class=\"fa fa-pencil\"><\/i>"
                });
                $('.editable-td-text').editable({
                    emptytext: "<i class=\"fa fa-pencil\"><\/i>",
                    type: "text"
                });
                $('.editable-td-datetime').editable({
                    "type": "combodate",
                    "emptytext": "<i class=\"fa fa-pencil\"><\/i>",
                    "format": "YYYY-MM-DD HH:mm:ss",
                    "viewformat": "YYYY-MM-DD HH:mm:ss",
                    "template": "YYYY-MM-DD HH:mm:ss",
                    "combodate": {"maxYear": 2035}
                });
                $('.editable-td-date').editable({
                    "type": "combodate",
                    "emptytext": "<i class=\"fa fa-pencil\"><\/i>",
                    "format": "YYYY-MM-DD",
                    "viewformat": "YYYY-MM-DD",
                    "template": "YYYY-MM-DD",
                    "combodate": {"maxYear": 2035}
                });
                $('.editable-td-year').editable({
                    "type": "combodate",
                    "emptytext": "<i class=\"fa fa-pencil\"><\/i>",
                    "format": "YYYY",
                    "viewformat": "YYYY",
                    "template": "YYYY",
                    "combodate": {"maxYear": 2035}
                });
                $('.editable-td-month').editable({
                    "type": "combodate",
                    "emptytext": "<i class=\"fa fa-pencil\"><\/i>",
                    "format": "MM",
                    "viewformat": "MM",
                    "template": "MM",
                    "combodate": {"maxYear": 2035}
                });
                $('.editable-td-day').editable({
                    "type": "combodate",
                    "emptytext": "<i class=\"fa fa-pencil\"><\/i>",
                    "format": "DD",
                    "viewformat": "DD",
                    "template": "DD",
                    "combodate": {"maxYear": 2035}
                });
                $('.editable-td-textarea').editable({
                    "type": "textarea",
                    "rows": 10,
                    "emptytext": "<i class=\"fa fa-pencil\"><\/i>"
                });
                $(".info_edit_switch").bootstrapSwitch({
                    onSwitchChange: function (event, state) {
                        let obejct = $(event.target);
                        let val = "";
                        if (state) {
                            val = obejct.closest('.bootstrap-switch').next().val();
                        } else {
                            val = obejct.closest('.bootstrap-switch').next().next().val()
                        }
                        $.ajax({
                            method: 'post',
                            url: obejct.data("updateurl"),
                            data: {
                                name: obejct.data("field"),
                                value: val,
                                pk: obejct.data("pk")
                            },
                            success: function (data) {
                                if (typeof (data) === "string") {
                                    data = JSON.parse(data);
                                }
                                if (data.code !== 200) {
                                    swal(data.msg, '', 'error');
                                }
                            }
                        });
                    }
                })
            });

            {{renderRowDataJS "" .ActionJs}}
        </script>
        <style>
            table tbody tr td {
                word-wrap: break-word;
                word-break: break-all;
            }
        </style>
    {{end}}
{{end}}`,"components/tabs":`{{define "tabs"}}
<div class="nav-tabs-custom">
    <ul class="nav nav-tabs">
        {{range $key, $data := .Data}}
            {{if eq $key 0}}
                <li class="active"><a href="#tab_{{$key}}" data-toggle="tab" aria-expanded="true">{{index $data "title"}}</a></li>
            {{else}}
                <li><a href="#tab_{{$key}}" data-toggle="tab" aria-expanded="true">{{index $data "title"}}</a></li>
            {{end}}
        {{end}}
    </ul>
    <div class="tab-content">
        {{range $key, $data := .Data}}
            {{if eq $key 0}}
                <div class="tab-pane active" id="tab_{{$key}}">
                {{index $data "content"}}
                </div>
            {{else}}
                <div class="tab-pane" id="tab_{{$key}}">
                {{index $data "content"}}
                </div>
            {{end}}
        {{end}}
    </div>
</div>
{{end}}`,"components/tree-header":`{{define "tree-header"}}
<div class="btn-group">
    <a class="btn btn-primary btn-sm tree-model-tree-tools" data-action="expand">
        <i class="fa fa-plus-square-o"></i>&nbsp;{{lang "expand"}}
    </a>
    <a class="btn btn-primary btn-sm tree-model-tree-tools" data-action="collapse">
        <i class="fa fa-minus-square-o"></i>&nbsp;{{lang "collapse"}}
    </a>
</div>

<div class="btn-group">
    <a class="btn btn-info btn-sm  tree-model-save"><i class="fa fa-save"></i>&nbsp;{{lang "save"}}</a>
</div>

<div class="btn-group">
    <a class="btn btn-warning btn-sm tree-model-refresh"><i class="fa fa-refresh"></i>&nbsp;{{lang "refresh"}}</a>
</div>
<div class="btn-group">
</div>
{{end}}`,"components/tree":`{{define "tree"}}
    <div class="dd" id="tree-model">
        {{$EditUrl := .EditUrl}}
        {{$UrlPrefix := .UrlPrefix}}
        <ol class="dd-list">
            {{range $key, $list := .Tree}}
                <li class="dd-item" data-id='{{$list.ID}}'>
                    <div class="dd-handle">
                        {{if eq $list.Url ""}}
                            <i class="fa {{$list.Icon}}"></i>&nbsp;<strong>{{$list.Name}}</strong>&nbsp;&nbsp;&nbsp;<a
                                    href="{{$list.Url}}" class="dd-nodrag">{{$list.Url}}</a>
                        {{else if eq $list.Url "/"}}
                            <i class="fa {{$list.Icon}}"></i>&nbsp;<strong>{{$list.Name}}</strong>&nbsp;&nbsp;&nbsp;<a
                                    href="{{$UrlPrefix}}" class="dd-nodrag">{{$UrlPrefix}}</a>
                        {{else if (isLinkUrl $list.Url)}}
                            <i class="fa {{$list.Icon}}"></i>&nbsp;<strong>{{$list.Name}}</strong>&nbsp;&nbsp;&nbsp;<a
                                    href="{{$list.Url}}" class="dd-nodrag">{{$list.Url}}</a>
                        {{else}}
                            <i class="fa {{$list.Icon}}"></i>&nbsp;<strong>{{$list.Name}}</strong>&nbsp;&nbsp;&nbsp;<a
                                    href="{{$UrlPrefix}}{{$list.Url}}" class="dd-nodrag">{{$UrlPrefix}}{{$list.Url}}</a>
                        {{end}}
                        <span class="pull-right dd-nodrag">
                <a href="{{$EditUrl}}?id={{$list.ID}}"><i class="fa fa-edit"></i></a>
                <a href="javascript:void(0);" data-id="{{$list.ID}}" class="tree_branch_delete"><i class="fa fa-trash"></i></a>
            </span>
                    </div>
                    {{if gt (len $list.ChildrenList) 0}}
                        <ol class="dd-list">
                            {{range $key, $item := $list.ChildrenList}}
                                <li class="dd-item" data-id='{{$item.ID}}'>
                                    <div class="dd-handle">
                                        {{if eq $item.Url ""}}
                                            <i class="fa {{$item.Icon}}"></i>&nbsp;<strong>{{$item.Name}}</strong>&nbsp;&nbsp;&nbsp;<a
                                                    href="{{$item.Url}}" class="dd-nodrag">{{$item.Url}}</a>
                                        {{else if eq $item.Url "/"}}
                                            <i class="fa {{$item.Icon}}"></i>&nbsp;<strong>{{$item.Name}}</strong>&nbsp;&nbsp;&nbsp;<a
                                                    href="{{$UrlPrefix}}" class="dd-nodrag">{{$UrlPrefix}}</a>
                                        {{else if (isLinkUrl $item.Url)}}
                                            <i class="fa {{$item.Icon}}"></i>&nbsp;<strong>{{$item.Name}}</strong>&nbsp;&nbsp;&nbsp;<a
                                                    href="{{$item.Url}}" class="dd-nodrag">{{$item.Url}}</a>
                                        {{else}}
                                            <i class="fa {{$item.Icon}}"></i>&nbsp;<strong>{{$item.Name}}</strong>&nbsp;&nbsp;&nbsp;<a
                                                    href="{{$UrlPrefix}}{{$item.Url}}" class="dd-nodrag">{{$UrlPrefix}}{{$item.Url}}</a>
                                        {{end}}
                                        <span class="pull-right dd-nodrag">
                                <a href="{{$EditUrl}}?id={{$item.ID}}"><i class="fa fa-edit"></i></a>
                                <a href="javascript:void(0);" data-id="{{$item.ID}}" class="tree_branch_delete"><i class="fa fa-trash"></i></a>
                            </span>
                                    </div>
                                </li>
                            {{end}}
                        </ol>
                    {{end}}
                </li>
            {{end}}
        </ol>
    </div>
    <script data-exec-on-popstate="">
        $(function () {
            $('#tree-model').nestable([]);
            $('.tree_branch_delete').click(function () {
                let id = $(this).data('id');
                swal({
                        title: {{lang "are you sure to delete"}} + "?",
                        type: "warning",
                        showCancelButton: true,
                        confirmButtonColor: "#DD6B55",
                        confirmButtonText: {{lang "confirm"}},
                        closeOnConfirm: false,
                        cancelButtonText: {{lang "cancel"}}
                    },
                    function () {
                        $.ajax({
                            method: 'post',
                            url: {{.DeleteUrl}} + '?id=' + id,
                            data: {},
                            success: function (data) {
                                $.pjax.reload('#pjax-container');
                                if (data.code === 200) {
                                    swal(data.msg, '', {{lang "success"}});
                                } else {
                                    swal(data.msg, '', {{lang "error"}});
                                }
                            }
                        });
                    });
            });
            $('.tree-model-save').click(function () {
                let serialize = $('#tree-model').nestable('serialize');
                $.post({{.OrderUrl}}, {
                        _order: JSON.stringify(serialize)
                    },
                    function (data) {
                        $.pjax.reload('#pjax-container');
                        toastr.success('Save succeeded !');
                    });
            });
            $('.tree-model-refresh').click(function () {
                $.pjax.reload('#pjax-container');
                toastr.success(toastMsg);
            });
            $('.tree-model-tree-tools').on('click', function (e) {
                let target = $(e.target),
                    action = target.data('action');
                if (action === 'expand') {
                    $('.dd').nestable('expandAll');
                }
                if (action === 'collapse') {
                    $('.dd').nestable('collapseAll');
                }
            });
            $(".parent_id").select2({"allowClear": true, "placeholder": "Parent"});
            $(".roles").select2({"allowClear": true, "placeholder": "Roles"});
        });
    </script>
{{end}}`,"content":`{{define "content"}}
    <script>
        $('a.new-tab-link').on('click', function () {
            listenerForAddNavTab($(this).attr('href'), $(this).attr('data-title'))
        });
    </script>
    <script src="{{link .CdnUrl .UrlPrefix "/assets/dist/js/content.min.js"}}"></script>
    {{.AssetsList}}
    {{if lang .Panel.Title}}
        <section class="content-header">
            <h1>
                {{lang .Panel.Title}}
                <small>{{lang .Panel.Description}}</small>
            </h1>
            <ol class="breadcrumb" style="margin-right: 30px;">
                <li><a href="{{.IndexUrl}}"><i class="fa fa-dashboard"></i> {{lang "home"}}</a></li>
                {{.Menu.FormatPath}}
            </ol>
        </section>
    {{end}}

    <!-- Main content -->
    <section class="content">
        {{.Panel.Content}}
    </section>
{{end}}`,"control_panel":`{{define "control_panel"}}
    <div class="control-sidebar-bg" style="position: fixed; height: auto;"></div>
    <aside class="control-sidebar control-sidebar-dark control-sidebar-open"
           style="position: fixed; max-height: 100%; overflow: auto;">
        <ul class="nav nav-tabs nav-justified control-sidebar-tabs">
            <li class="active"><a href="#control-sidebar-setting-tab" data-toggle="tab" aria-expanded="true"><i
                            class="fa fa-wrench"></i></a></li>
            <li><a href="#control-sidebar-home-tab" data-toggle="tab"><i class="fa fa-home"></i></a></li>
            <li><a href="#control-sidebar-settings-tab" data-toggle="tab"><i class="fa fa-gears"></i></a></li>
        </ul>
        <div class="tab-content">
            <div class="tab-pane active" id="control-sidebar-setting-tab">
                <h4 class="control-sidebar-heading">{{lang "layout"}}</h4>
                <div class="form-group"><label class="control-sidebar-subheading"><input type="checkbox"
                                                                                         data-layout="fixed"
                                                                                         class="pull-right"
                                                                                         checked="checked"> 固定布局</label>
                    <p>盒子模型和固定布局不能同时启作用</p></div>
                <div class="form-group"><label class="control-sidebar-subheading"><input type="checkbox"
                                                                                         data-layout="layout-boxed"
                                                                                         class="pull-right">
                        盒子布局</label>
                    <p>盒子布局最大宽度将被限定为1250px</p></div>
                <div class="form-group"><label class="control-sidebar-subheading"><input type="checkbox"
                                                                                         data-layout="sidebar-collapse"
                                                                                         class="pull-right">
                        切换菜单栏</label>
                    <p>切换菜单栏的展示或收起</p></div>
                <div class="form-group"><label class="control-sidebar-subheading"><input type="checkbox"
                                                                                         data-enable="expandOnHover"
                                                                                         class="pull-right">
                        菜单栏自动展开</label>
                    <p>鼠标移到菜单栏自动展开</p></div>
                <div class="form-group"><label class="control-sidebar-subheading"><input type="checkbox"
                                                                                         data-menu="show-submenu"
                                                                                         class="pull-right">
                        显示菜单栏子菜单</label>
                    <p>菜单栏子菜单将始终显示</p></div>
                <div class="form-group"><label class="control-sidebar-subheading"><input type="checkbox"
                                                                                         data-menu="disable-top-badge"
                                                                                         class="pull-right"
                                                                                         checked="checked">
                        禁用顶部彩色小角标</label>
                    <p>左边菜单栏的彩色小角标不受影响</p></div>
                <div class="form-group"><label class="control-sidebar-subheading"><input type="checkbox"
                                                                                         data-controlsidebar="control-sidebar-open"
                                                                                         class="pull-right">
                        切换右侧操作栏</label>
                    <p>切换右侧操作栏覆盖或独占</p></div>
                <div class="form-group"><label class="control-sidebar-subheading"><input type="checkbox"
                                                                                         data-sidebarskin="toggle"
                                                                                         class="pull-right">
                        切换右侧操作栏背景</label>
                    <p>将右侧操作栏背景亮色或深色切换</p></div>
                <h4 class="control-sidebar-heading">{{lang "skin"}}</h4>
                <ul class="list-unstyled clearfix skin-list">
                    <li><a href="javascript:;" data-skin="skin-blue" style="" class="clearfix full-opacity-hover">
                            <div>
                                <span style="display:block; width: 20%; float: left; height: 7px; background: #367fa9;"></span><span
                                        class="bg-light-blue"
                                        style="display:block; width: 80%; float: left; height: 7px;"></span></div>
                            <div>
                                <span style="display:block; width: 20%; float: left; height: 20px; background: #222d32;"></span><span
                                        style="display:block; width: 80%; float: left; height: 20px; background: #f4f5f7;"></span>
                            </div>
                        </a>
                        <p class="text-center no-margin">Blue</p></li>
                    <li><a href="javascript:;" data-skin="skin-white" class="clearfix full-opacity-hover">
                            <div style="box-shadow: 0 0 2px rgba(0,0,0,0.1)" class="clearfix"><span
                                        style="display:block; width: 20%; float: left; height: 7px; background: #fefefe;"></span><span
                                        style="display:block; width: 80%; float: left; height: 7px; background: #fefefe;"></span>
                            </div>
                            <div>
                                <span style="display:block; width: 20%; float: left; height: 20px; background: #222;"></span><span
                                        style="display:block; width: 80%; float: left; height: 20px; background: #f4f5f7;"></span>
                            </div>
                        </a>
                        <p class="text-center no-margin">White</p></li>
                    <li><a href="javascript:;" data-skin="skin-purple" class="clearfix full-opacity-hover">
                            <div><span style="display:block; width: 20%; float: left; height: 7px;"
                                       class="bg-purple-active"></span><span class="bg-purple"
                                                                             style="display:block; width: 80%; float: left; height: 7px;"></span>
                            </div>
                            <div>
                                <span style="display:block; width: 20%; float: left; height: 20px; background: #222d32;"></span><span
                                        style="display:block; width: 80%; float: left; height: 20px; background: #f4f5f7;"></span>
                            </div>
                        </a>
                        <p class="text-center no-margin">Purple</p></li>
                    <li><a href="javascript:;" data-skin="skin-green" class="clearfix full-opacity-hover">
                            <div><span style="display:block; width: 20%; float: left; height: 7px;"
                                       class="bg-green-active"></span><span class="bg-green"
                                                                            style="display:block; width: 80%; float: left; height: 7px;"></span>
                            </div>
                            <div>
                                <span style="display:block; width: 20%; float: left; height: 20px; background: #222d32;"></span><span
                                        style="display:block; width: 80%; float: left; height: 20px; background: #f4f5f7;"></span>
                            </div>
                        </a>
                        <p class="text-center no-margin">Green</p></li>
                    <li><a href="javascript:;" data-skin="skin-red" class="clearfix full-opacity-hover">
                            <div><span style="display:block; width: 20%; float: left; height: 7px;"
                                       class="bg-red-active"></span><span class="bg-red"
                                                                          style="display:block; width: 80%; float: left; height: 7px;"></span>
                            </div>
                            <div>
                                <span style="display:block; width: 20%; float: left; height: 20px; background: #222d32;"></span><span
                                        style="display:block; width: 80%; float: left; height: 20px; background: #f4f5f7;"></span>
                            </div>
                        </a>
                        <p class="text-center no-margin">Red</p></li>
                    <li><a href="javascript:;" data-skin="skin-yellow" class="clearfix full-opacity-hover">
                            <div><span style="display:block; width: 20%; float: left; height: 7px;"
                                       class="bg-yellow-active"></span><span class="bg-yellow"
                                                                             style="display:block; width: 80%; float: left; height: 7px;"></span>
                            </div>
                            <div>
                                <span style="display:block; width: 20%; float: left; height: 20px; background: #222d32;"></span><span
                                        style="display:block; width: 80%; float: left; height: 20px; background: #f4f5f7;"></span>
                            </div>
                        </a>
                        <p class="text-center no-margin">Yellow</p></li>
                    <li><a href="javascript:;" data-skin="skin-blue-light" class="clearfix full-opacity-hover">
                            <div>
                                <span style="display:block; width: 20%; float: left; height: 7px; background: #367fa9;"></span><span
                                        class="bg-light-blue"
                                        style="display:block; width: 80%; float: left; height: 7px;"></span></div>
                            <div>
                                <span style="display:block; width: 20%; float: left; height: 20px; background: #f9fafc;"></span><span
                                        style="display:block; width: 80%; float: left; height: 20px; background: #f4f5f7;"></span>
                            </div>
                        </a>
                        <p class="text-center no-margin" style="font-size: 12px">Blue Light</p></li>
                    <li><a href="javascript:;" data-skin="skin-white-light" class="clearfix full-opacity-hover">
                            <div style="box-shadow: 0 0 2px rgba(0,0,0,0.1)" class="clearfix"><span
                                        style="display:block; width: 20%; float: left; height: 7px; background: #fefefe;"></span><span
                                        style="display:block; width: 80%; float: left; height: 7px; background: #fefefe;"></span>
                            </div>
                            <div>
                                <span style="display:block; width: 20%; float: left; height: 20px; background: #f9fafc;"></span><span
                                        style="display:block; width: 80%; float: left; height: 20px; background: #f4f5f7;"></span>
                            </div>
                        </a>
                        <p class="text-center no-margin" style="font-size: 12px">White Light</p></li>
                    <li><a href="javascript:;" data-skin="skin-purple-light" class="clearfix full-opacity-hover">
                            <div><span style="display:block; width: 20%; float: left; height: 7px;"
                                       class="bg-purple-active"></span><span class="bg-purple"
                                                                             style="display:block; width: 80%; float: left; height: 7px;"></span>
                            </div>
                            <div>
                                <span style="display:block; width: 20%; float: left; height: 20px; background: #f9fafc;"></span><span
                                        style="display:block; width: 80%; float: left; height: 20px; background: #f4f5f7;"></span>
                            </div>
                        </a>
                        <p class="text-center no-margin" style="font-size: 12px">Purple Light</p></li>
                    <li><a href="javascript:;" data-skin="skin-green-light" class="clearfix full-opacity-hover">
                            <div><span style="display:block; width: 20%; float: left; height: 7px;"
                                       class="bg-green-active"></span><span class="bg-green"
                                                                            style="display:block; width: 80%; float: left; height: 7px;"></span>
                            </div>
                            <div>
                                <span style="display:block; width: 20%; float: left; height: 20px; background: #f9fafc;"></span><span
                                        style="display:block; width: 80%; float: left; height: 20px; background: #f4f5f7;"></span>
                            </div>
                        </a>
                        <p class="text-center no-margin" style="font-size: 12px">Green Light</p></li>
                    <li><a href="javascript:;" data-skin="skin-red-light" class="clearfix full-opacity-hover">
                            <div><span style="display:block; width: 20%; float: left; height: 7px;"
                                       class="bg-red-active"></span><span class="bg-red"
                                                                          style="display:block; width: 80%; float: left; height: 7px;"></span>
                            </div>
                            <div>
                                <span style="display:block; width: 20%; float: left; height: 20px; background: #f9fafc;"></span><span
                                        style="display:block; width: 80%; float: left; height: 20px; background: #f4f5f7;"></span>
                            </div>
                        </a>
                        <p class="text-center no-margin" style="font-size: 12px">Red Light</p></li>
                    <li><a href="javascript:;" data-skin="skin-yellow-light" class="clearfix full-opacity-hover">
                            <div><span style="display:block; width: 20%; float: left; height: 7px;"
                                       class="bg-yellow-active"></span><span class="bg-yellow"
                                                                             style="display:block; width: 80%; float: left; height: 7px;"></span>
                            </div>
                            <div>
                                <span style="display:block; width: 20%; float: left; height: 20px; background: #f9fafc;"></span><span
                                        style="display:block; width: 80%; float: left; height: 20px; background: #f4f5f7;"></span>
                            </div>
                        </a>
                        <p class="text-center no-margin" style="font-size: 12px;">Yellow Light</p></li>
                </ul>
            </div>
            <div class="tab-pane" id="control-sidebar-home-tab">
                <h4 class="control-sidebar-heading">{{lang "home"}}</h4>
            </div>
            <div class="tab-pane" id="control-sidebar-settings-tab">
                <h4 class="control-sidebar-heading">{{lang "config"}}</h4>
            </div>
        </div>
    </aside>
{{end}}`,"footer":`{{define "footer"}}
    <footer class="main-footer">
        <div class="pull-right hidden-xs">
            <b>Version</b> {{.System.Version}}
        </div>
        <strong>Powered by <a href="https://github.com/GoAdminGroup/go-admin">GoAdmin</a>.</strong>
        {{.FooterInfo}}
    </footer>
{{end}}`,"head":`{{define "head"}}
    <head>
        <meta charset="utf-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <title>{{.Title}}</title>
        <!-- Tell the browser to be responsive to screen width -->
        <meta content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no" name="viewport">

        <link rel="stylesheet" href="{{link .CdnUrl .UrlPrefix "/assets/dist/css/all.min.css"}}">

        <!--[if lt IE 9]>
        <script src="{{link .CdnUrl .UrlPrefix "/assets/dist/js/html5shiv.min.js"}}"></script>
        <script src="{{link .CdnUrl .UrlPrefix "/assets/dist/js/respond.min.js"}}"></script>
        <![endif]-->

        <script src="{{link .CdnUrl .UrlPrefix "/assets/dist/js/all.min.js"}}"></script>

        {{.CustomHeadHtml}}
    </head>
{{end}}`,"header":`{{define "header"}}
    <header class="main-header">
        <!-- Logo -->
        <a href="{{.UrlPrefix}}" class="logo">
            <!-- mini logo for sidebar mini 50x50 pixels -->
            <span class="logo-mini">{{.MiniLogo}}</span>
            <!-- logo for regular state and mobile devices -->
            <span class="logo-lg">{{.Logo}}</span>
        </a>
        <!-- Header Navbar: style can be found in header.less -->
        <nav class="navbar navbar-static-top">
            <div id="firstnav">
                <a href="#" class="sidebar-toggle" data-toggle="offcanvas" role="button">
                    <span class="sr-only">Toggle navigation</span>
                </a>

                <ul class="nav nav-tabs nav-addtabs" style="max-width: 60%">
                </ul>

                {{ template "admin_panel" . }}
            </div>
        </nav>
    </header>
{{end}}`,"js":`{{define "js"}}
    <script src="{{link .CdnUrl .UrlPrefix "/assets/dist/js/all_2.min.js"}}"></script>
{{end}}`,"layout":`{{define "layout"}}

    <!DOCTYPE html>
    <html>

    {{ template "head" . }}

    <body class="hold-transition {{.ColorScheme}} sidebar-mini">
    <div class="wrapper">

        {{ template "header" . }}

        {{ template "sidebar" . }}

        <div class="content-wrapper" id="pjax-container">

            {{ template "content" . }}

        </div>

        {{ template "footer" . }}

    </div>

    {{ template "js" . }}

    </body>
    {{.CustomFootHtml}}
    </html>

{{end}}
`,"menu":`{{define "menu"}}
    <ul class="sidebar-menu" data-widget="tree">
        {{$UrlPrefix := .UrlPrefix}}
        {{range $key, $list := .Menu.List }}
            {{if eq (len $list.ChildrenList) 0}}
                {{if $list.Header}}
                    <li class="header" data-rel="external">{{$list.Header}}</li>
                {{end}}
                <li class='{{$list.Active}}'>
                    {{if eq $list.Url "/"}}
                    <a href='{{$UrlPrefix}}'>
                    {{else if isLinkUrl $list.Url}}
                    <a href='{{$list.Url}}'>
                    {{else}}
                    <a href='{{$UrlPrefix}}{{$list.Url}}'>
                    {{end}}
                        <i class="fa {{$list.Icon}}"></i><span> {{$list.Name}}</span>
                        <span class="pull-right-container"><!-- <small class="label pull-right bg-green">new</small> --></span>
                    </a>
                </li>
            {{else}}
                <li class="treeview {{$list.Active}}">
                    <a href="#">
                        <i class="fa {{$list.Icon}}"></i><span> {{$list.Name}}</span>
                        <span class="pull-right-container">
                        <i class="fa fa-angle-left pull-right"></i>
                    </span>
                    </a>
                    <ul class="treeview-menu">
                        {{range $key2, $item := $list.ChildrenList}}
                            <li><a href="{{$UrlPrefix}}{{$item.Url}}"><i class="fa {{$item.Icon}}"></i> {{$item.Name}}
                                </a></li>
                        {{end}}
                    </ul>
                </li>
            {{end}}
        {{end}}
    </ul>
{{end}}`,"sidebar":`{{define "sidebar"}}
    <!-- Left side column. contains the logo and sidebar -->
    <aside class="main-sidebar">
        <!-- sidebar: style can be found in sidebar.less -->
        <section class="sidebar">
            <!-- Sidebar user panel -->
            <div class="user-panel">
                <div class="pull-left image">
                    {{if eq .User.Avatar ""}}
                        <img src="{{.UrlPrefix}}/assets/dist/img/avatar04.png" class="img-circle" alt="User Image">
                    {{else}}
                        <img src="{{.User.Avatar}}" class="img-circle" alt="User Image">
                    {{end}}
                </div>
                <div class="pull-left info">
                    {{.User.Name}}
                    <a href="#"><i class="fa fa-circle text-success"></i> {{lang "online"}}</a>
                </div>
            </div>
            <!-- search form -->
            <!-- <form action="#" method="get" class="sidebar-form">
                <div class="input-group">
                    <input type="text" name="q" class="form-control" placeholder="Search...">
                    <span class="input-group-btn">
                <button type="submit" name="search" id="search-btn" class="btn btn-flat"><i class="fa fa-search"></i>
                </button>
              </span>
                </div>
            </form>-->
            <!-- /.search form -->

            {{ template "menu" . }}

        </section>
        <!-- /.sidebar -->
    </aside>
{{end}}`,}