<template>
    <div>
        <div>blog</div>
        <div>
            {{ id }}
        </div>
    </div>
</template>

<script>
import axios from 'axios';

export default {
    data() {
        return {
            id: 0,
            blog: {}
        }
    },
    methods: {
        getBlog() {
            axios.get('api/blog/id/' + this.id).then(res => {
                if (res.data.code === 0) {
                    this.blog = res.data.data
                } else if (res.data.code === 1003) {
                    this.$message({
                        message: '文章不存在',
                        type: 'error'
                    })
                }
            }).catch(() => {
                this.$message({
                    message: '服务器错误',
                    type: 'error'
                })
            })
        }
    },
    mounted() {
        this.id = this.$route.query.id
    }
}
</script>