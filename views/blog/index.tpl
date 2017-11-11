{{template "blog/header.tpl" .}}
<script type="text/javascript">document.title = '博客 - {{config "String" "title" ""}}';</script>

    <body>
        <div class="container">
            <div class="left-col">
                <div class="intrude-less">
                    <header id="header" class="inner">
                        <div class="profilepic"><a href="/"></a></div>
                        <h1><a href="/blog">All...</a></h1>
                        <p class="subtitle">{{config "String" "say" ""}}</p>
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
                    <div id="content" class="inner">
                        {{range .DocList}}
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
                                    <p>
                                        <script type="text/javascript">document.write({{substr .Markdown 0 100}}.replace(/\n/g,"<br/>"));</script>
                                    &   <br/>...
                                    </p>
                                    <p><a href="/blog/doc/{{.DocumentId}}" class="more-link">继续阅读 »</a></p>
                                </div>
                            </article>
                        {{end}}
                        <nav class="page-navi">
                            <a href="?page={{.PagePrev}}" class="prev">« 上一页</a>
                            <a href="?page={{.PageNext}}" class="next">下一页 »</a>
                        </nav>                        
                    </div>
                </div>
                {{template "blog/foot.tpl" .}}
            </div>
        </div>
    </body>
</html>