let emtt = {
	tasks: {},
	on(key, func) {
		if (this.tasks.hasOwnProperty(key)) {
			this.tasks[key].push(func);
		} else {
			this.tasks[key] = [func];
		}
	},
	emit(key, ...args) {
		this.tasks[key].forEach(fn => {
			fn(...args)
		});
	},
	off(key, func) {
		if (this.tasks.hasOwnProperty(key)) {
			if (func) {
				this.tasks[key] = this.tasks[key].filters((item) => {
					return func !== item;
				});
			} else {
				this.tasks[key] = [];
			}
		}
	},
};


// emtt.on('keep',function(...val){
//     console.log('val',val);
// })
// emtt.emit('keep',2,3,4,5);
export default emtt