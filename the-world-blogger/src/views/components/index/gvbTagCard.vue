<template>
  <gvb-card title="标签云" class="tags_cloud_card" top20>
    <div>
      <span @click="allBlogs" style="margin-left: 40px; color: #2184fc; cursor: pointer">全部</span>
    </div>
    <ul class="tags_cloud_ul">
      <li :class="{active: item.id === $store.state.tagId}" v-for="(item, index) in this.tags" :key="index">
        <span @click="checkTag(item.id)">
          {{ item.name }}
        </span>
      </li>
    </ul>
  </gvb-card>
</template>

<script>
import gvbCard from "@/views/components/gvbCard";
import axios from "axios";

export default {
  name: "gvbTagCard",
  components: {
    gvbCard
  },
  data() {
    return {
      tags: []
    }
  },
  props: {
    tagNums: Number
  },
  methods: {
    getTags() {
      axios.get("/api/tag/search?limit=" + this.$props.tagNums + "&offset=0").then(res => {
        if (res.data.code === 0) {
          this.tags = res.data.data.tag
        } else {
          this.$message({
            message: '获取标签失败',
            type: 'warning'
          })
        }
      })
    },
    checkTag(tagId) {
      this.$store.state.tagId = tagId
    },
    allBlogs() {
      this.$store.state.tagId = 0
    }
  },
  created() {
    this.getTags()
  }
}
</script>

<style lang="scss">
.tags_cloud_ul {
  display: flex;
  flex-wrap: wrap;

  li {
    width: 33.33%;
    height: 40px;
    display: flex;
    align-items: center;
    justify-content: center;
    color: var(--h2);


    span {
      cursor: pointer;
      &:hover{
        color: #2184fc;
      }
    }

    &.active{
      color: #2184fc;
    }
  }

  li:nth-child(6n+1), li:nth-child(6n+2), li:nth-child(6n+3) {
    background-color: #f0eeee;
  }

  li:nth-child(3n+1), li:nth-child(3n+2) {
    border-right: 1px solid var(--bg_darke);
  }

}
</style>