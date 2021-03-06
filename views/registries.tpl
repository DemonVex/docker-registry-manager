{{template "base/base.html" .}}
{{define "body"}}
{{template "new_registry.tpl" .}}
  <div class="right-content-container">
    <div class="header">
      <ol class="breadcrumb">
        <li><a href="/">Home</a></li>
        <li><a href="/registries">Registries</a></li>
      </ol>
    </div>

    <div class = "content-block-empty">
      <div class="col-lg-12">
        <ul class="boxes">
          {{range $key, $registry := .registries}}
          <li>
          <a href="/registries/{{$registry.Host}}/repositories">
            <div class="white-bg box col-lg-4 col-md-6 col-sm-12 col-xs-12">
              <div class="col-lg-12">
                <div class="box-container">
                  <div class="box-header">
                    <h2>{{$registry.Host}}<small> {{$registry.IP}}</small></h2>
                  </div>
                  <div class="box-body">
                    <div class="info">
                      <div class="info-container">
                        <h3 class="info-metric">{{len $registry.Repositories}}</h3>
                        <small>
                          {{ $repoCount := len $registry.Repositories }} {{ if eq $repoCount 1 }} Repository {{else}} Repositories {{ end }}
                        </small>
                      </div>
                    </div>
                    <div class="info">
                      <div class="info-container">
                        <h3 class="info-metric">{{$registry.TagCount}}</h3>
                        <small>
                          {{ $tagCount := $registry.TagCount }} {{ if eq $tagCount 1 }} Tag {{else}} Tags {{ end }}
                        </small>
                      </div>
                    </div>
                    <div class="info">
                      <div class="info-container">
                        <h3 class="info-metric">{{$registry.DiskSize}}</h3>
                        <small>Total Size</small>
                      </div>
                    </div>
                  </div>
                  <div class="box-footer">
                    <span class="label label-success text-capitalize">{{$registry.Status}}</span>
                  </div>
                </div>
              </div>
            </div>
          </a>
          </li>
          {{end}}
          <li>
            <div class="well-box box col-lg-4 col-md-6 col-sm-12 col-xs-12">
              <div class="col-lg-12">
                <div class="box-container">
                  <div type= "button" class="add-new" data-toggle="modal" data-target="#new-registry-modal">
                    <i class="fa fa-plus add-new-icon"></i>
                  </div>
                </div>
              </div>
            </div>
          </li>
        </ul>
      </div>
    </div>
  </div>

{{end}}
