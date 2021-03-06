// app.ts

import {formatTimeTwo} from "./utils/util";

let resolveUserInfo: any;

App<IAppOption>({
    globalData: {
        userInfo: new Promise((resolve) => {
            resolveUserInfo = resolve;
        })
    },
    onLaunch() {
        let that = this
        wx.login({
            success: res => {
                wx.request({
                    url: 'http://localhost:9002/v1/auth/login',
                    method: 'POST',
                    data: {code: res.code},
                    timeout: 6000,
                    success: (response) => {
                        let data = response.data as any;
                        that.setTokenToStorage(data.token, parseInt(data.aging) + Date.parse(Date()) / 1000)
                    },
                    fail: console.log
                });
            },
        })

        console.log("时间", formatTimeTwo((new Date("1994-01-13")).getTime(), "Y-M-D"))
    },
    parseUserInfo(e: WechatMiniprogram.UserInfo) {
        resolveUserInfo(e)
    },
    setTokenToStorage(token: string, aging: number): void {
        wx.setStorageSync("tokenInfo", JSON.stringify({token: token, aging: aging}))
    }
})