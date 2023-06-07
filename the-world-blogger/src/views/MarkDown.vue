<template>
  <div class="gvb_md_view">
    <div>
      <el-input type="textarea" autosize placeholder="请输入题目" v-model="title">
      </el-input>
    </div>
    <div style="margin: 10px 0;"></div>
    <div>
      <el-input type="textarea" placeholder="请输入文章简介" v-model="profile" maxlength="50" show-word-limit>
      </el-input>
    </div>
    <div>
      <el-upload
          class="upload-demo"
          :file-list="coverList"
          :before-upload="beforeAvatarUpload"
          :on-remove="handleRemove"
          list-type="picture">
        <el-button size="small" type="primary">点击上传文章封面</el-button>
      </el-upload>
    </div>
    <div style="margin: 10px 0;"></div>
    <mavon-editor v-model="content" ref="md" @change="change" @imgAdd="$imgAdd" @imgDel="$imgDel"
                  style="min-height: 600px"/>
    <button @click="submit">提交</button>
  </div>
</template>

<script>
import axios from 'axios';
import {mavonEditor} from 'mavon-editor'
import 'mavon-editor/dist/css/index.css'

export default {
  // 注册
  components: {
    mavonEditor,
  },
  data() {
    return {
      content: '', // 输入的markdown
      html: '',    // 及时转的html
      title: '',
      profile: '',
      cover: '',
      img_file: {},
      coverList: [],
    }
  },
  methods: {
    // 所有操作都会被解析重新渲染
    change(value, render) {
      // render 为 markdown 解析后的结果[html]
      this.html = render;
    },
    // 提交
    submit() {
      let blogData = {}
      blogData.title = this.title
      blogData.content = this.content
      blogData.profile = this.profile
      blogData.cover = this.cover
      axios.post('/api/blog/add', blogData).then(res => {
        if (res.data.code === 0) {
          // todo 成功后跳转文章列表
          this.$message({
            message: '提交成功',
            type: 'success'
          });
          this.$router.push('/')
        } else if (res.data.code === 1006) {
          this.$message({
            message: '请先登录',
            type: 'warning'
          })
          this.$router.push('/login')
        }
      }).catch(() => {
      })
    },
    $imgAdd(pos, $file) {
      var formdata = new FormData()
      formdata.append('image', $file)
      axios({
        url: '/api/common/upload',
        method: 'post',
        data: formdata,
        headers: {'Content-Type': 'multipart/form-data'},
      }).then(res => {
        // 第二步.将返回的url替换到文本原位置![...](0) -> ![...](url)
        /**
         * $vm 指为mavonEditor实例，可以通过如下两种方式获取
         * 1. 通过引入对象获取: `import {mavonEditor} from ...` 等方式引入后，`$vm`为`mavonEditor`
         * 2. 通过$refs获取: html声明ref : `<mavon-editor ref=md ></mavon-editor>，`$vm`为 `this.$refs.md`
         */
        this.$refs.md.$img2Url(pos, res.data.data);
      })
    },
    // handleAvatarSuccess(res, file) {
    //
    // },
    beforeAvatarUpload(file) {
      var formdata = new FormData()
      formdata.append('image', file)
      axios({
        url: '/api/common/upload',
        method: 'post',
        data: formdata,
        headers: {'Content-Type': 'multipart/form-data'},
      }).then(res => {
        // 第二步.将返回的url替换到文本原位置![...](0) -> ![...](url)
        /**
         * $vm 指为mavonEditor实例，可以通过如下两种方式获取
         * 1. 通过引入对象获取: `import {mavonEditor} from ...` 等方式引入后，`$vm`为`mavonEditor`
         * 2. 通过$refs获取: html声明ref : `<mavon-editor ref=md ></mavon-editor>，`$vm`为 `this.$refs.md`
         */
        if (this.cover !== '') {
          this.coverList[0].name = file.name
          this.coverList[0].url = res.data.data
        } else {
          this.coverList.push({
            name: file.name,
            url: res.data.data
          })
        }
        this.cover = res.data.data
      })
    },
    handleRemove() {
      this.cover = ''
      this.coverList.pop()
    }
  },
  mounted() {
    if (localStorage.getItem('loginStatus') === 'false') {
      this.$message({
        message: '请先登录',
        type: 'warning'
      });
      this.$router.push('/login')
    }
    axios.defaults.headers.common['Authorization'] = localStorage.getItem("token")
  }
}
</script>

<style lang="scss">
.gvb_md_view {
  img {
    font-size: 87%;
  }
}
</style>