<template>
    <div>
        <el-card>
            <el-form ref="form" :model="form" label-width="80px">
              <el-form-item label="论文名称">
                <el-input v-model="form.name" placeholder="请输入论文名称" :disabled="activeTab === 'updatePaper'"></el-input>
              </el-form-item>
              <el-form-item label="作者姓名">
                <el-input v-model="form.author" placeholder="!为了有较好的显示效果，请只输入第一(二)作者"></el-input>
              </el-form-item>
              <el-form-item label="会议(期刊)">
                <el-input v-model="form.inst" placeholder="请输入会议(期刊)名称"></el-input>
              </el-form-item>
              <el-form-item label="代码地址">
                <el-input v-model="form.code" placeholder="请输入项目代码地址: 例如 https://www.github.com/BUAA/Demo"></el-input>
              </el-form-item>
              <el-form-item label="类型" v-model="form.type">
                <el-radio-group v-model="form.type">
                  <el-radio label="c">会议论文</el-radio>
                  <el-radio label="j">期刊论文</el-radio>
                  <el-radio label="b">毕业设计</el-radio>
                </el-radio-group>
              </el-form-item>
              <el-form-item label="录用时间">
                <el-col :span="11">
                  <el-date-picker type="month" placeholder="选择日期" v-model="form.date" format="yyyy 年 MM 月" style="width: 100%;"></el-date-picker>
                </el-col>
              </el-form-item>
<!--                TODO 添加论文标签 最佳论文 or 优秀毕业论文-->
<!--                TODO 删除已经上传的ppt或pdf-->
              <el-form-item label="上传材料">
                  <el-upload
                      ref="upload"
                      drag
                      :action="baseUrl+'/api/paper/uploadfiles'"
                      :data="paper_name"
                      :on-success="cleanFile"
                      :on-error="clear"
                      multiple
                      :auto-upload="false"
                  >
                      <i class="el-icon-upload"></i>
                      <div class="el-upload__text">将文件拖到此处，或<em>点击上传</em></div>
                    </el-upload>
                    <el-alert
                        style="width: 725px"
                        title="上传文件的时候请上传.pdf .doc .docx(论文)或者.ppt .pptx .key(PPT/KeyNote)结尾的文件, 两类文件可以同时上传"
                        type="warning">
                    </el-alert>
              </el-form-item>
              <el-form-item>
                <el-button type="primary" @click="onSubmit" v-if="activeTab==='addPaper'">立即创建</el-button>
                <el-button type="primary" @click="onSubmit" v-else>立即更新</el-button>
                <el-button @click="clear">取消</el-button>
              </el-form-item>
            </el-form>
        </el-card>
    </div>
</template>

<script>
    import axios from "axios"
    export default {
        name: "AddPaper",
        data() {
          return {
            fileList: [],
            form: {
                name: "", author: "", inst: "",
                code: "", type: "", date: "",
            }
          }
        },
        computed:{
          activeTab(){
              return this.$store.state.app.activeTab
          },
          baseUrl(){
              return this.$store.state.app.baseUrl
          },
          paper_name(){
            return {
                name: this.form.name
            }
          }
        },
        watch:{
          activeTab: function (val, Oldval) {
              if(val === "updatePaper"){
                  if(!this.$store.state.app.editPaper){
                      this.$store.commit("setActiveTab", "viewPaper");
                      this.$notify({ title: '注意', message: '请先选择要更新的论文'+this.form.name, type: 'warning'});
                      return
                  }
                  this.form = this.$store.state.app.editPaper;
                  delete this.form.article;
                  delete this.form.create;
                  delete this.form.ppt;
              } else {
                  this.form =  {
                      name: "", author: "", inst: "",
                      code: "", type: "", date: "",
                  }
              }

          }
        },
        methods:{
            cleanFile(response, file, fileList){
                if(this.activeTab==="addPaper"){
                   this.$notify({ title: '文件上传成功', message: '成功添加论文'+this.form.name, type: 'success'});
                } else {
                  this.$notify({ title: '文件上传成功', message: '成功更新论文'+this.form.name, type: 'success'});
                }
                this.clear();
                setTimeout(()=>{
                    this.$store.commit("setActiveTab", "viewPaper")
                }, 1000)
            },
            clear(){
                this.$refs.upload.clearFiles();
                this.form = {
                    name: "", author: "", inst: "",
                    code: "", type: "", date: "",
                };
            },
            format(date){
                Date.prototype.Format = function (fmt) {
                  var o = {
                    "y+": this.getFullYear(),
                    "M+": this.getMonth() + 1,
                    "d+": this.getDate(),
                    "h+": this.getHours(),
                    "m+": this.getMinutes(),
                    "s+": this.getSeconds(),
                    "q+": Math.floor((this.getMonth() + 3) / 3),
                    "S+": this.getMilliseconds()
                  };
                  for (var k in o) {
                    if (new RegExp("(" + k + ")").test(fmt)){
                      if(k === "y+"){
                        fmt = fmt.replace(RegExp.$1, ("" + o[k]).substr(4 - RegExp.$1.length));
                      }
                      else if(k ==="S+"){
                        var lens = RegExp.$1.length;
                        lens = lens === 1?3:lens;
                        fmt = fmt.replace(RegExp.$1, ("00" + o[k]).substr(("" + o[k]).length - 1,lens));
                      }
                      else{
                        fmt = fmt.replace(RegExp.$1, (RegExp.$1.length === 1) ? (o[k]) : (("00" + o[k]).substr(("" + o[k]).length)));
                      }
                    }
                  }
                  return fmt;
                };
                return date.Format("yyyy-MM-dd hh:mm:ss")
            },
            onSubmit(){
                if(this.form.name===""){
                    this.$notify({
                      title: '注意',
                      message: '请填写论文名称',
                      type: 'error'
                    });
                    return
                }
                if(this.form.type===""){
                    this.$notify({
                      title: '注意',
                      message: '请选择论文类型',
                      type: 'error'
                    });
                    return
                }
                if(this.form.date instanceof Date) {
                    this.form.date = this.format(this.form.date);
                }
                axios.post(this.baseUrl+"/admin/paper?token="+localStorage.getItem("labac"), this.form)
                    .then(res => {
                        if(res.data.code === 1){
                          this.$notify({ title: '论文信息更新成功', message: '成功更新论文信息'+this.form.name, type: 'success'});
                          this.$refs.upload.submit()
                        }
                        else {
                            this.$notify({ title: '注意', message: '创建失败', type: 'error'});
                        }
                    })
                    .catch(err => {
                        this.clear();
                        this.$notify({ title: '注意', message: err, type: 'error'});
                    });
            }
        }
    }
</script>

<style scoped>
</style>