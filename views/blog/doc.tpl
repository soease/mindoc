{{template "blog/header.tpl" .}}
<script type="text/javascript">document.title = '{{.Doc.DocumentName}} - {{.Member.Account}} - {{config "String" "title" ""}}';</script>
    <body>
        <div class="container">
            <div class="left-col">
                <div class="intrude-less">
                    <header id="header" class="inner">
                        {{with .Member}}
                        <div style="background-image:url({{.Avatar}})" class="profilepic" ></div>
                        <h1><a href="/blog/user/{{.MemberId}}">{{.Account}}</a></h1>
                        <p class="subtitle">{{.Description}}</p>
                        {{end}}
                        <nav id="main-nav">
                            <ul>
                                <li class="on" ><a href="/blog"><span>首页</span></a></li>
                                {{range .Lists}}                                
                                    <li><a href="/blog/book/{{.BookId}}">{{.BookName}}</a></li>
                                {{end}}
                                <li ><a href="/blog/about"><span>关于</span></a></li>
                                <li ><a href="/blog/search"><span>搜索</span></a></li>
                            </ul>
                        </nav>
                    </header>
                </div>
            </div>
            <div class="mid-col">
                <div class="mid-col-container">
                    <div id="content" class="inner" style="margin:10px 10px 10px 15px;">
                        {{with .Doc}}
                            <article class="post post-list">
                                <div class="meta">
                                    <div class="date">
                                        {{dateformat .CreateTime "2006-01-02 15:04:05"}}
                                    </div>
                                    <div class="date">
                                        {{dateformat .ModifyTime "2006-01-02 15:04:05"}}
                                    </div>
                                </div>
                                <h1 class="title"><a href="/blog/doc/{{.DocumentId}}">
                                    {{.DocumentName}}
                                </a></h1>
                                <div class="entry-content">
                                    <p>{{str2html .Release}}</p>
                                </div>
                            </article>
                        {{end}}
                        <nav class="page-navi">
                            {{with .DocPrev}}<a href="/blog/doc/{{.DocumentId}}" class="prev">« {{.DocumentName}}</a>{{end}}
                            {{with .DocNext}}<a href="/blog/doc/{{.DocumentId}}" class="next">{{.DocumentName}} »</a>{{end}}
                        </nav>                        
                    </div>
                </div>
                {{template "blog/foot.tpl" .}}
            </div>
        </div>
    </body>
</html>