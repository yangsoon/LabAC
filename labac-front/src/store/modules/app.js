const app = {
  state :{
      activeMenu: "paper",
      activeTab: "viewPaper",
      editPaper: null,
      baseUrl: "",
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
