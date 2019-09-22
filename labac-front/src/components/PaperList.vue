<template>
<div>
    <el-container>
      <el-aside width="200px" style="height: 100%">
        <el-menu @select="filterPaper">
          <el-menu-item index="all">
            <template slot="title">
                <span style="margin-right: 20px">
                    <svg class="ali-icon" aria-hidden="true">
                        <use xlink:href="#ali-icon-notebook"></use>
                    </svg>
                </span>
                全部论文
            </template>
          </el-menu-item>
          <el-menu-item :index="year" v-for="(year, index) in yearList" :key="index">
            <template slot="title">
                <span style="margin-right: 20px">
                    <svg class="ali-icon" aria-hidden="true">
                        <use xlink:href="#ali-icon-shijian"></use>
                    </svg>
                </span>
                {{year}}
            </template>
          </el-menu-item>
        </el-menu>
      </el-aside>
        <el-main style="height: 800px">
            <el-timeline>
                <el-timeline-item :timestamp="item.date" placement="top" v-for="(item, index) in items" :key="index">
                    <el-card>
                      <div slot="header" class="clearfix">
                          <span>
                            <span style="margin-right: 20px">
                                <svg class="ali-icon" aria-hidden="true">
                                    <use xlink:href="#ali-icon-lunwentimu"></use>
                                </svg>
                            </span>
                              <strong>{{item.name}}</strong>
                          </span>
                      </div>
                      <div>
                          <div class="item">
                               <a :href="baseUrl+'/api/paper/download/'+item.name+'/article'" class="link" v-show="item.article">
                                   <svg class="ali-icon" aria-hidden="true" v-if="item.articleSuffix==='pdf'">
                                        <use xlink:href="#ali-icon-pdf"></use>
                                   </svg>
                                   <svg class="ali-icon" aria-hidden="true" v-else>
                                        <use xlink:href="#ali-icon-docx"></use>
                                   </svg>
                                   原文
                               </a>
                          </div>
                          <div class="item" v-show="item.ppt">
                            <a :href="baseUrl+'/api/paper/download/'+item.name+'/ppt'" class="link">
                                <div v-if="item.pptSuffix==='key'">
                                    <svg class="ali-icon" aria-hidden="true">
                                        <use xlink:href="#ali-icon-Keynotejiangyan"></use>
                                    </svg>
                                    KeyNote
                                </div>
                                <div v-else>
                                    <svg class="ali-icon" aria-hidden="true">
                                        <use xlink:href="#ali-icon-ppttubiao"></use>
                                    </svg>
                                    PPT
                                </div>
                            </a>
                          </div>
                          <div class="item" v-show="item.code">
                              <a :href="item.code" target="_blank" class="link">
                                  <svg class="ali-icon" aria-hidden="true">
                                    <use xlink:href="#ali-icon-github"></use>
                                  </svg>
                                  代码
                              </a>
                          </div>
                      </div>
                      <div>
                          <span>
                            <span style="margin-right: 20px">
                                <svg class="ali-icon" aria-hidden="true">
                                    <use xlink:href="#ali-icon-huabanfuben"></use>
                                </svg>
                            </span>
                              <strong>作者:</strong>&nbsp; {{item.author}}
                          </span> &nbsp;&nbsp;
                          <span v-show="item.type==='c'">
                               <span style="margin-right: 20px">
                                    <svg class="ali-icon" aria-hidden="true">
                                        <use xlink:href="#ali-icon-C"></use>
                                    </svg>
                                </span>
                              <strong>会议论文:</strong>&nbsp; {{item.inst}}
                          </span>
                          <span v-show="item.type==='j'">
                               <span style="margin-right: 20px">
                                    <svg class="ali-icon" aria-hidden="true">
                                        <use xlink:href="#ali-icon-J"></use>
                                    </svg>
                                </span>
                              <strong>期刊论文:</strong>&nbsp; {{item.inst}}
                          </span>
                          <span v-show="item.type==='b'">
                               <span style="margin-right: 20px">
                                    <svg class="ali-icon" aria-hidden="true">
                                        <use xlink:href="#ali-icon-nanbiyesheng"></use>
                                    </svg>
                                </span>
                              <strong>毕业论文:</strong>&nbsp; {{item.inst}}
                          </span>
                      </div>
                    </el-card>
                </el-timeline-item>
            </el-timeline>
        </el-main>
    </el-container>
</div>
</template>

<script>
    import axios from "axios"
    export default {
        name: "PaperList",
        data(){
            return{
                items: [],
                source: [],
                yearList: [],
            }
        },
        computed:{
            baseUrl(){
              return this.$store.state.app.baseUrl
            }
        },
        created() {
            axios.get(this.baseUrl+"/api/paper").then(res=>{
                let temp = res.data.data;
                temp.sort(this.cmp("date"));
                this.items = temp;
                this.source = temp;
                this.yearList = this.getYearList(temp);
            })
        },
        methods:{
            filterPaper(year, _){
                if(year === 'all') {
                    this.items = this.source;
                    return
                }
                let temp = [];
                for(let i = 0; i < this.source.length; i++){
                    if(this.source[i].year === year){
                        temp.push(this.source[i])
                    }
                }
                this.items = temp
            },
            getYearList(src){
                let years = new Set();
                for(let i=0; i < src.length; i++){
                    let year = src[i]["date"].split("-")[0];
                    years.add(year);
                    this.source[i]['year'] = year
                }
                return Array.from(years)
            },
            cmp(p){
                return function (a, b) {
                    let x = parseInt(a[p].replace('-',''));
                    let y = parseInt(b[p].replace('-',''));
                    return y - x
                }
            }
        }
    }
</script>

<style scoped>
    .item {
        float: right;
        padding: 3px 0;
        margin: 0 10px;
    }
    a {
        text-decoration:none;
        color:inherit;
    }
    .link{
        height:30px;
        line-height:30px;
    }
    .ali-icon {
      width: 1.5em;
      height: 1.5em;
      vertical-align: -0.15em;
      fill: currentColor;
      overflow: hidden;
    }
</style>
