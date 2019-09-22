<template>
    <div>
        <el-container>
            <el-aside width="200px" style="height: 100%">
                <el-menu default-active="paper">
                  <el-menu-item index="paper">
                    <i class="el-icon-edit-outline"></i>论文编辑
                  </el-menu-item>
                  <el-menu-item index="home">
                    <svg class="ali-icon" aria-hidden="true" style="margin-right: 10px">
                      <use xlink:href="#ali-icon-college"></use>
                    </svg>
                    首页编辑
                  </el-menu-item>
                </el-menu>
            </el-aside>
            <el-main>
              <el-tabs v-model="activeTab">
                <el-tab-pane label="管理论文" name="viewPaper">
                    <PaperTable></PaperTable>
                </el-tab-pane>
                <el-tab-pane label="添加论文" name="addPaper">
                    <AddPaper></AddPaper>
                </el-tab-pane>
                <el-tab-pane label="更新论文" name="updatePaper">
                    <AddPaper></AddPaper>
                </el-tab-pane>
              </el-tabs>
            </el-main>
        </el-container>
    </div>
</template>

<script>
  import AddPaper from "./AddPaper"
  import PaperTable from "./PaperTable"
  export default {
    components:{
      AddPaper,
      PaperTable
    },
    computed: {
        activeTab:{
            set: function (value) {
                this.$store.commit("setActiveTab", value);
            },
            get: function () {
                return this.$store.state.app.activeTab
            }
        }
    }
  }
</script>

<style scoped>
.ali-icon {
  width: 1.2em;
  height: 1.2em;
  vertical-align: -0.15em;
  fill: currentColor;
  overflow: hidden;
}
</style>