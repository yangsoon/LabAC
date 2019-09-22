<template>
    <div>
        <el-table
            v-loading="loading"
            stripe
            height="700"
            width="100%"
            :data="tableData"
            style="width: 100%">
            <el-table-column
              label="添加论文时间"
              width="200">
              <template slot-scope="scope">
                <i class="el-icon-time"></i>
                <span style="margin-left: 10px">{{scope.row.create}}</span>
              </template>
            </el-table-column>
            <el-table-column
              label="论文名称"
            >
                <template slot-scope="scope">{{scope.row.name}}</template>
            </el-table-column>
            <el-table-column
              label="论文类型"
              width="120"
              align="center"
            >
                <template slot-scope="scope">{{scope.row.inst}}</template>
            </el-table-column>
            <el-table-column
              label="原文提交"
              width="80"
              align="center"
            >
                <template slot-scope="scope">
                    <div>
                        <i v-if="scope.row.article === ''" class="el-icon-error" style="color: #F56C6C"></i>
                        <i v-else class="el-icon-success"></i>
                    </div>
                </template>
            </el-table-column>
            <el-table-column
              label="PPT提交"
              width="80"
              align="center"
            >
                <template slot-scope="scope">
                    <div>
                        <i v-if="scope.row.ppt === ''" class="el-icon-error" style="color: #F56C6C"></i>
                        <i v-else class="el-icon-success"></i>
                    </div>
                </template>
            </el-table-column>
            <el-table-column label="操作">
              <template slot-scope="scope">
                <el-button
                  size="mini"
                  @click="handleEdit(scope.$index, scope.row)">编辑</el-button>
                <el-button
                  size="mini"
                  type="danger"
                  @click="handleDelete(scope.$index, scope.row)">删除</el-button>
              </template>
            </el-table-column>
        </el-table>
    </div>
</template>

<script>
import axios from "axios"
export default {
    name: "PaperTable",
    data() {
      return {
        tableData: null,
        loading: true,
      }
    },
    created(){
        axios.get(this.baseUrl+'/api/paper').then(res => {
            this.tableData = res.data.data;
            this.loading = false
        })
    },
    computed:{
      baseUrl(){
          return this.$store.state.app.baseUrl
      },
      activeTab(){
          return this.$store.state.app.activeTab
      }
    },
    watch:{
      activeTab: function (val, Oldval) {
          if(val === "viewPaper"){
            axios.get(this.baseUrl+'/api/paper').then(res => {
                this.tableData = res.data.data;
            })
          }
      }
    },
    methods: {
      handleEdit(index, row) {
        this.$store.commit("setEditPaper", row);
        this.$store.commit("setActiveTab", "updatePaper")
      },
      handleDelete(index, row) {
        if(this.$store.state.app.editPaper && row.name === this.$store.state.app.editPaper.name){
            this.$store.commit("setEditPaper", null)
        }
        axios.delete(this.baseUrl+"/admin/paper/"+row.name+"?token="+localStorage.getItem("labac")).then(res => {
            if(res.data.code === 1){
                this.$notify({ title: '删除成功', type: 'success'});
                this.tableData.splice(index,1)
            }
        }).catch(err => {
            this.$notify({ title: '删除失败', type: 'error', message: "非法操作"});
        })
      }
    }
}
</script>

<style scoped>

</style>