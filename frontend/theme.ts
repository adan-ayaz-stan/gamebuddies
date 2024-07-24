import type { CustomThemeConfig } from '@skeletonlabs/tw-plugin';

export const myCustomTheme: CustomThemeConfig = {
	name: 'my-custom-theme',
	properties: {
		// =~= Theme Properties =~=
		'--theme-font-family-base': `system-ui`,
		'--theme-font-family-heading': `system-ui`,
		'--theme-font-color-base': '0 0 0',
		'--theme-font-color-dark': '255 255 255',
		'--theme-rounded-base': '6px',
		'--theme-rounded-container': '6px',
		'--theme-border-base': '2px',
		// =~= Theme On-X Colors =~=
		'--on-primary': '255 255 255',
		'--on-secondary': '0 0 0',
		'--on-tertiary': '255 255 255',
		'--on-success': '0 0 0',
		'--on-warning': '0 0 0',
		'--on-error': '255 255 255',
		'--on-surface': '255 255 255',
		// =~= Theme Colors  =~=
		// primary | #FF5A5F
		'--color-primary-50': '255 230 231', // #ffe6e7
		'--color-primary-100': '255 222 223', // #ffdedf
		'--color-primary-200': '255 214 215', // #ffd6d7
		'--color-primary-300': '255 189 191', // #ffbdbf
		'--color-primary-400': '255 140 143', // #ff8c8f
		'--color-primary-500': '255 90 95', // #FF5A5F
		'--color-primary-600': '230 81 86', // #e65156
		'--color-primary-700': '191 68 71', // #bf4447
		'--color-primary-800': '153 54 57', // #993639
		'--color-primary-900': '125 44 47', // #7d2c2f
		// secondary | #FFFFFF
		'--color-secondary-50': '255 255 255', // #ffffff
		'--color-secondary-100': '255 255 255', // #ffffff
		'--color-secondary-200': '255 255 255', // #ffffff
		'--color-secondary-300': '255 255 255', // #ffffff
		'--color-secondary-400': '255 255 255', // #ffffff
		'--color-secondary-500': '255 255 255', // #FFFFFF
		'--color-secondary-600': '230 230 230', // #e6e6e6
		'--color-secondary-700': '191 191 191', // #bfbfbf
		'--color-secondary-800': '153 153 153', // #999999
		'--color-secondary-900': '125 125 125', // #7d7d7d
		// tertiary | #FF0008
		'--color-tertiary-50': '255 217 218', // #ffd9da
		'--color-tertiary-100': '255 204 206', // #ffccce
		'--color-tertiary-200': '255 191 193', // #ffbfc1
		'--color-tertiary-300': '255 153 156', // #ff999c
		'--color-tertiary-400': '255 77 82', // #ff4d52
		'--color-tertiary-500': '255 0 8', // #FF0008
		'--color-tertiary-600': '230 0 7', // #e60007
		'--color-tertiary-700': '191 0 6', // #bf0006
		'--color-tertiary-800': '153 0 5', // #990005
		'--color-tertiary-900': '125 0 4', // #7d0004
		// success | #84cc16
		'--color-success-50': '237 247 220', // #edf7dc
		'--color-success-100': '230 245 208', // #e6f5d0
		'--color-success-200': '224 242 197', // #e0f2c5
		'--color-success-300': '206 235 162', // #ceeba2
		'--color-success-400': '169 219 92', // #a9db5c
		'--color-success-500': '132 204 22', // #84cc16
		'--color-success-600': '119 184 20', // #77b814
		'--color-success-700': '99 153 17', // #639911
		'--color-success-800': '79 122 13', // #4f7a0d
		'--color-success-900': '65 100 11', // #41640b
		// warning | #EAB308
		'--color-warning-50': '252 244 218', // #fcf4da
		'--color-warning-100': '251 240 206', // #fbf0ce
		'--color-warning-200': '250 236 193', // #faecc1
		'--color-warning-300': '247 225 156', // #f7e19c
		'--color-warning-400': '240 202 82', // #f0ca52
		'--color-warning-500': '234 179 8', // #EAB308
		'--color-warning-600': '211 161 7', // #d3a107
		'--color-warning-700': '176 134 6', // #b08606
		'--color-warning-800': '140 107 5', // #8c6b05
		'--color-warning-900': '115 88 4', // #735804
		// error | #D41976
		'--color-error-50': '249 221 234', // #f9ddea
		'--color-error-100': '246 209 228', // #f6d1e4
		'--color-error-200': '244 198 221', // #f4c6dd
		'--color-error-300': '238 163 200', // #eea3c8
		'--color-error-400': '225 94 159', // #e15e9f
		'--color-error-500': '212 25 118', // #D41976
		'--color-error-600': '191 23 106', // #bf176a
		'--color-error-700': '159 19 89', // #9f1359
		'--color-error-800': '127 15 71', // #7f0f47
		'--color-error-900': '104 12 58', // #680c3a
		// surface | #3E0505
		'--color-surface-50': '226 218 218', // #e2dada
		'--color-surface-100': '216 205 205', // #d8cdcd
		'--color-surface-200': '207 193 193', // #cfc1c1
		'--color-surface-300': '178 155 155', // #b29b9b
		'--color-surface-400': '120 80 80', // #785050
		'--color-surface-500': '62 5 5', // #3E0505
		'--color-surface-600': '56 5 5', // #380505
		'--color-surface-700': '47 4 4', // #2f0404
		'--color-surface-800': '37 3 3', // #250303
		'--color-surface-900': '30 2 2' // #1e0202
	}
};
