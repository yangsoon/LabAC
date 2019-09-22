const app = {
  state :{
      activeMenu: "paper",
      activeTab: "viewPaper",
      editPaper: null,
      baseUrl: "http://127.0.0.1:8000",
  },
  mutations:{
      setActiveMenu(state, menu){
          state.activeMenu = menu
      },
      setActiveTab(state, tab){
          state.activeTab = tab
      },
      setEditPaper(state, paper){
          state.editPaper = paper
      }
  }
};
export default app;
