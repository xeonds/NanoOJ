<template>
  <div>
    <h1>排行榜</h1>
    <RankList></RankList>
  </div>
</template>

<script>
import RankList from '../../components/RankList.vue';

export default {
  name: "RanklistView",
  components: {
    RankList
  }
};
</script>
