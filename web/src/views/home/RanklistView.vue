<template>
  <div>
    <h1>排行榜</h1>
    <RankList :ranks="ranks"></RankList>
  </div>
</template>

<script>
import RankList from '../../components/RankList.vue';
import { mapGetters } from 'vuex';

export default {
  name: "RanklistView",
  components: {
    RankList
  },
  computed: {
    ...mapGetters({ ranks: "getRanks" }),
  },
};
</script>
