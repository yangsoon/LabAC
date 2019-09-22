<template>
<el-container>
  <el-header style="padding: 0px">
    <el-menu
      mode="horizontal"
      background-color="#545c64"
      text-color="#fff"
      active-text-color="#ffd04b"
      @select="handlerRouter"
      :default-active="activeMenu"
    >
      <el-menu-item index="home">
        <svg class="ali-icon" aria-hidden="true" style="margin-right: 10px">
          <use xlink:href="#ali-icon-ziyuan"></use>
        </svg>
        <span slot="title">首页</span>
      </el-menu-item>
      <el-menu-item index="paper">
        <svg class="ali-icon" aria-hidden="true" style="margin-right: 10px">
          <use xlink:href="#ali-icon-folder"></use>
        </svg>
        <span slot="title">论文</span>
      </el-menu-item>
      <el-menu-item index="edit" v-if="showEdit">
        <svg class="ali-icon" aria-hidden="true" style="margin-right: 10px">
          <use xlink:href="#ali-icon-edit"></use>
        </svg>
        <span slot="title">编辑</span>
      </el-menu-item>
      <el-menu-item index="login" style="float: right" @click="showLogin = true" v-if="!showEdit" >
        <svg class="ali-icon" aria-hidden="true" style="margin-right: 2px; ">
          <use xlink:href="#ali-icon-paper-airplane"></use>
        </svg>
        <span slot="title">登录</span>
      </el-menu-item>
      <el-menu-item index="logout" style="float: right" @click="logout" v-if="showEdit">
        <svg class="ali-icon" aria-hidden="true" style="margin-right: 10px;">
          <use xlink:href="#ali-icon-tuichu"></use>
        </svg>
        <span slot="title">注销</span>
      </el-menu-item>
      <el-dialog
        title="登录"
        :visible.sync="showLogin"
        width="22%"
        center>
         <el-input v-model="username" style="margin: 5% 0">
          <template slot="prepend">
            <svg class="ali-icon" aria-hidden="true">
              <use xlink:href="#ali-icon-yonghu"></use>
            </svg>
            账号
          </template>
        </el-input>
        <el-input v-model="password" type="password">
          <template slot="prepend" >
            <svg class="ali-icon" aria-hidden="true">
              <use xlink:href="#ali-icon-key"></use>
            </svg>
            密码
          </template>
        </el-input>
        <span slot="footer" class="dialog-footer" style="margin-bottom: 5%">
          <el-button type="primary" @click="Login">登录</el-button>
          <el-button @click="cancelLogin">取 消</el-button>
        </span>
        <el-alert style="margin-top: 5%" type="error" v-show="showAlert">{{message}}</el-alert>
      </el-dialog>
    </el-menu>
  </el-header>
  <el-main style="height: 100%; border: 1px solid #eee">
    <router-view></router-view>
  </el-main>
  <el-footer>
    <div class="copyright"><p>© 2019 <a href="http://github.com/yangsoon">yangsoon</a></p></div>
  </el-footer>
</el-container>
</template>

<script>
    import axios from "axios"
    export default {
        name: "AchieveMent",
        data(){
          return{
            showLogin: false,
            showAlert: false,
            showEdit: false,
            username: "",
            password: "",
            message: ""
          }
        },
        created(){
          if(localStorage.getItem("labac")){
              this.showEdit = true;
          }
        },
        computed:{
          baseUrl(){
              return this.$store.state.app.baseUrl
          },
          activeMenu: {
            get(){
              return this.$store.state.app.activeMenu;
            },
            set(value){
              this.$store.commit("setActiveMenu", value);
            }
          }
        },
        methods:{
          Login(){
            if(this.username ==='' || this.password === ''){
              this.message = "请填写完整用户名和密码";
              this.showAlert = true;
              setTimeout(()=>{
                this.message = "";
                this.showAlert = false
              }, 1000);
            } else{
              axios.get(this.baseUrl+"/login", {
                params: {
                  username: this.username,
                  password: this.password
                }
              }).then(res => {
                this.message = "";
                this.showAlert = false;
                if(res.data.code === 1){
                  this.showEdit = true;
                  localStorage.setItem("labac", res.data.token);
                  this.showLogin = false;
                  this.$router.push({name: 'edit'});
                  this.$notify({
                    title: "成功",
                    message: "欢迎管理员登录",
                    type: "success"
                  })
                } else {
                  this.message = "用户名或密码错误";
                  this.showAlert = true;
                }
              })
            }

          },
          cancelLogin(){
            this.showLogin = false
          },
          logout(){
            localStorage.removeItem("labac");
            this.$router.push({name: 'paper'});
            this.showEdit = false;
            this.$notify({
              title: "注销成功",
              message: "再见👋",
              type: "success"
            })
          },
          handlerRouter(key, keyPath){
            if(key==='login'||key==='logout'){
              return
            }
            this.$router.push({
              name: key
            })
          }
        }
    }
</script>

<style scoped>
.copyright {
    margin: 4em 0;
    border-top: 1px solid #ddd;
    text-align: center;
}
a {
  text-decoration:none;
  color:inherit;
}
.ali-icon {
  width: 1.2em;
  height: 1.2em;
  vertical-align: -0.15em;
  fill: currentColor;
  overflow: hidden;
}

</style>