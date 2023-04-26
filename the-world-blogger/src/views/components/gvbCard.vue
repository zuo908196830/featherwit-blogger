<template>
<div class="gvb_card_view">
  <div class="title">
    <h2>{{ this.$props.title }}</h2>
    <a v-if="this.$props.link_name !== ''" :href="this.$props.link">{{ this.$props.link_name }}</a>
  </div>
  <div :class="{body: true, top20: this.$props.top20}">
    <slot></slot>
  </div>
</div>
</template>

<script>

export default {
  name: "gvuCard",
  props: {
    title: String,
    link_name: String,
    link: String,
    top20: Boolean
  }
}
</script>

<style lang="scss">
.gvb_card_view {
  background-color: white;

  .title {
    padding: 20px 20px 10px 20px;
    display: flex;
    justify-content: space-between;
    align-items: center;
    border-radius: 5px 5px 0 0;
    background-color: white;

    h2 {
      color: #333333;
      font-size: 20px;
      font-weight: 600;
    }
    a {
      color: #555;
    }
  }

  .body {
    border-radius: 0 0 5px 5px;
    background-color: white;
    margin-top: 1px;
    padding: 10px 20px 20px 20px;
  }

  .body.top20 {
    padding: 20px;
  }
}
</style>