export default {
  getRankById: (state) => (id) => {
    return state.ranks.find((rank) => rank.id == id);
  },
  getRanks: (state) => {
    return state.ranks;
  },
};
