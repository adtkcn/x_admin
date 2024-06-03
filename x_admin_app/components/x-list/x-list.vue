<template >
    <view class="x-list" :style="{marginBottom:bottom+'px'}">
 
        <view class="x-list-left" v-if="icon">
            <uv-image width="24px" height="24px" :src="icon" mode="widthFix"></uv-image>
        </view>

        <view class="x-list-content">
            <view v-if="title">{{title}}</view>
            <text v-if="des">{{des}}</text>
        </view>
        <view class="x-list-right" v-if="btnText">
            <slot name="right">
                <button  :open-type="openType" @click="btnClick">{{btnText}}</button>
            </slot>            
        </view>
    </view>
</template>
<script>
export default {
    props:{
        icon:{
            default:"",
            type:String
        },
        title:{
            default:"",
            type:String
        },
        des:{
            default:"",
            type:String
        },
        btnText:{
            default:"",
            type:String
        },
        openType:{
            default:"",
            type:String
        },
        bottom:{
            default:"0",
            type:[String,Number]
        },
        url:{
            default:"",
            type:String
        },
    },
	data() {
		return {
			
		}
	},
    computed:{
    },
	methods: {
        btnClick(){
            if(this.url){
                this.$toPath(this.url, {
					// id: item.id
				})
                return
            }
            this.$emit('btnClick')
        },
	}
}
</script>
<style lang="scss">
.x-list {
    display: flex;
    // justify-content: center;
    align-items: center;
    background-color: white;
 
    padding: 10px 10px;
    border-radius: 8px;
    // margin-bottom: 10px;

    .x-list-content {
        font-size: 16px;
        flex: 1;
        padding: 0 10px;

        text {
            line-height: 1;
            font-size: 13px;
            color: #AAAAAA;
        }
    }

    .x-list-right {
        button {
            margin: 0;
            padding: 6px 8px;
            line-height: 1;
            border: 0;
            font-size: 12px;
            background-color: rgba(111, 148, 205, 1);
            color: white;
        }
    }
}
</style>