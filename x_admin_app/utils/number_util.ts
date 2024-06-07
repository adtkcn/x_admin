import BigNumber from "bignumber.js";

/**
 * 减法a-b
 * @param {*} a 
 * @param {*} b 
 */
export function sub(a, b) {
	return BigNumber(a).minus(b);
}
/**
 * 加法a+b
 * @param {*} a 
 * @param {*} b 
 */
export function add(a, b) {
    return BigNumber(a).plus(b);
}