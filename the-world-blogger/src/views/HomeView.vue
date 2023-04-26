<template>
  <div class="base_view index_view">
    <gvb-nav/>
    <div class="gvb_banner">
    </div>
    <div class="gvb_base_container">
      <div class="gvb_inner_container gvb_index_main">
        <div class="left">
          <gvb-blog-list style="margin-top: 20px"></gvb-blog-list>
        </div>
        <div class="right">
          <gvb-tag-card :tag-nums="15"></gvb-tag-card>
          <gvb-hot-card style="margin-top: 20px"></gvb-hot-card>
        </div>
      </div>
    </div>
    <gvb-footer></gvb-footer>
  </div>
</template>

<script>
import gvbNav from './components/gvbNav'
import gvbTagCard from "@/views/components/index/gvbTagCard";
import gvbFooter from "@/views/components/gvbFooter";
import gvbHotCard from "@/views/components/index/gvbHotCard";
import gvbBlogList from "@/views/components/index/gvbBlogList";
import GvbBlogList from "@/views/components/index/gvbBlogList";

export default {
  components: {
    GvbBlogList,
    gvbNav,
    gvbTagCard,
    gvbFooter,
    gvbHotCard,
    gvbBlogList
  },
  data() {
    return {
    }
  },
  methods: {
  },
  created() {

  },
}
</script>

<style lang="scss">
a {
  text-decoration: none;
}

li {
  list-style: none;
}

.base_view {
  background-color: #f0eeee;

  .gvb_base_container {
    top: 120px;
    display: flex;
    justify-content: center;

    .gvb_inner_container {
      width: 1200px;
      margin-top: 20px;
    }
  }

  .gvb_banner {
    height: 50px;
    width: 100%;
    overflow: hidden;
  }
}

.index_view {
  .gvb_inner_container {
    display: flex;
    justify-content: space-between;
  }

  .gvb_index_main {
    .left {
      width: calc(100% - 416px);
    }

    .right {
      width: 396px;
    }
  }


}

.user_show {
  cursor: pointer;

  &:hover {
    color: cornflowerblue;
  }
}
</style>