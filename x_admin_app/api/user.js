import {request} from '@/utils/request';



export function login(data) {
	return request({
		url: '/system/login',
		method: 'POST',
		data
	})
}

export function getInfo(token) {
	return request({
		url: '/system/admin/self',
		method: 'GET',
		data: {
			token
		}
	})
}

export function logout() {
	return request({
		url: '/system/logout',
		method: 'GET'
	})
}

