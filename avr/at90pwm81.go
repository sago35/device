// Automatically generated file. DO NOT EDIT.
// Generated by gen-device-avr.go from AT90PWM81.atdf, see http://packs.download.atmel.com/

// +build avr,at90pwm81

// Device information for the AT90PWM81.
package avr

import (
	"runtime/interrupt"
	"runtime/volatile"
	"unsafe"
)

// Some information about this device.
const (
	DEVICE = "AT90PWM81"
	ARCH   = "AVR8"
	FAMILY = "megaAVR"
)

// Interrupts
const (
	IRQ_RESET         = 0  // External Pin, Power-on Reset, Brown-out Reset, Watchdog Reset and JTAG AVR Reset
	IRQ_PSC2_CAPT     = 1  // PSC2 Capture Event
	IRQ_PSC2_EC       = 2  // PSC2 End Cycle
	IRQ_PSC2_EEC      = 3  // PSC2 End Of Enhanced Cycle
	IRQ_PSC0_CAPT     = 4  // PSC0 Capture Event
	IRQ_PSC0_EC       = 5  // PSC0 End Cycle
	IRQ_PSC0_EEC      = 6  // PSC0 End Of Enhanced Cycle
	IRQ_ANALOG_COMP_1 = 7  // Analog Comparator 1
	IRQ_ANALOG_COMP_2 = 8  // Analog Comparator 2
	IRQ_ANALOG_COMP_3 = 9  // Analog Comparator 3
	IRQ_INT0          = 10 // External Interrupt Request 0
	IRQ_TIMER1_CAPT   = 11 // Timer/Counter1 Capture Event
	IRQ_TIMER1_OVF    = 12 // Timer/Counter1 Overflow
	IRQ_ADC           = 13 // ADC Conversion Complete
	IRQ_INT1          = 14 // External Interrupt Request 1
	IRQ_SPI_STC       = 15 // SPI Serial Transfer Complet
	IRQ_INT2          = 16 // External Interrupt Request 2
	IRQ_WDT           = 17 // Watchdog Timeout Interrupt
	IRQ_EE_READY      = 18 // EEPROM Ready
	IRQ_SPM_READY     = 19 // Store Program Memory Read
	IRQ_max           = 19 // Highest interrupt number on this device.
)

// Map interrupt numbers to function names.
// These aren't real calls, they're removed by the compiler.
var (
	_ = interrupt.Register(IRQ_RESET, "__vector_RESET")
	_ = interrupt.Register(IRQ_PSC2_CAPT, "__vector_PSC2_CAPT")
	_ = interrupt.Register(IRQ_PSC2_EC, "__vector_PSC2_EC")
	_ = interrupt.Register(IRQ_PSC2_EEC, "__vector_PSC2_EEC")
	_ = interrupt.Register(IRQ_PSC0_CAPT, "__vector_PSC0_CAPT")
	_ = interrupt.Register(IRQ_PSC0_EC, "__vector_PSC0_EC")
	_ = interrupt.Register(IRQ_PSC0_EEC, "__vector_PSC0_EEC")
	_ = interrupt.Register(IRQ_ANALOG_COMP_1, "__vector_ANALOG_COMP_1")
	_ = interrupt.Register(IRQ_ANALOG_COMP_2, "__vector_ANALOG_COMP_2")
	_ = interrupt.Register(IRQ_ANALOG_COMP_3, "__vector_ANALOG_COMP_3")
	_ = interrupt.Register(IRQ_INT0, "__vector_INT0")
	_ = interrupt.Register(IRQ_TIMER1_CAPT, "__vector_TIMER1_CAPT")
	_ = interrupt.Register(IRQ_TIMER1_OVF, "__vector_TIMER1_OVF")
	_ = interrupt.Register(IRQ_ADC, "__vector_ADC")
	_ = interrupt.Register(IRQ_INT1, "__vector_INT1")
	_ = interrupt.Register(IRQ_SPI_STC, "__vector_SPI_STC")
	_ = interrupt.Register(IRQ_INT2, "__vector_INT2")
	_ = interrupt.Register(IRQ_WDT, "__vector_WDT")
	_ = interrupt.Register(IRQ_EE_READY, "__vector_EE_READY")
	_ = interrupt.Register(IRQ_SPM_READY, "__vector_SPM_READY")
)

// Peripherals.
var (
	// Fuses
	EXTENDED = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2)))
	HIGH     = (*volatile.Register8)(unsafe.Pointer(uintptr(0x1)))
	LOW      = (*volatile.Register8)(unsafe.Pointer(uintptr(0x0)))

	// Lockbits
	LOCKBIT = (*volatile.Register8)(unsafe.Pointer(uintptr(0x0)))

	// I/O Port
	PORTB = (*volatile.Register8)(unsafe.Pointer(uintptr(0x25)))
	DDRB  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x24)))
	PINB  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x23)))
	PORTD = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2b)))
	DDRD  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2a)))
	PIND  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x29)))
	PORTE = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2e)))
	DDRE  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2d)))
	PINE  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2c)))

	// Digital-to-Analog Converter
	DACL  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x58)))
	DACH  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x59)))
	DACON = (*volatile.Register8)(unsafe.Pointer(uintptr(0x76)))

	// Serial Peripheral Interface
	SPCR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x37)))
	SPSR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x38)))
	SPDR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x56)))

	// Watchdog Timer
	WDTCSR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x82)))

	// External Interrupts
	EICRA = (*volatile.Register8)(unsafe.Pointer(uintptr(0x89)))
	EIMSK = (*volatile.Register8)(unsafe.Pointer(uintptr(0x41)))
	EIFR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x40)))

	// Analog-to-Digital Converter
	ADMUX   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x28)))
	ADCSRA  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x26)))
	ADCL    = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4c)))
	ADCH    = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4d)))
	ADCSRB  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x27)))
	DIDR0   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x77)))
	DIDR1   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x78)))
	AMP0CSR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x79)))

	// Analog Comparator
	AC3CON  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x7f)))
	AC1CON  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x7d)))
	AC2CON  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x7e)))
	ACSR    = (*volatile.Register8)(unsafe.Pointer(uintptr(0x20)))
	AC3ECON = (*volatile.Register8)(unsafe.Pointer(uintptr(0x7c)))
	AC2ECON = (*volatile.Register8)(unsafe.Pointer(uintptr(0x7b)))
	AC1ECON = (*volatile.Register8)(unsafe.Pointer(uintptr(0x7a)))

	// CPU Registers
	SREG    = (*volatile.Register8)(unsafe.Pointer(uintptr(0x5f)))
	SPL     = (*volatile.Register8)(unsafe.Pointer(uintptr(0x5d)))
	SPH     = (*volatile.Register8)(unsafe.Pointer(uintptr(0x5e)))
	MCUCR   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x55)))
	MCUSR   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x54)))
	OSCCAL  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x88)))
	CLKPR   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x83)))
	SMCR    = (*volatile.Register8)(unsafe.Pointer(uintptr(0x53)))
	GPIOR2  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3b)))
	GPIOR1  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3a)))
	GPIOR0  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x39)))
	PLLCSR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x87)))
	PRR     = (*volatile.Register8)(unsafe.Pointer(uintptr(0x86)))
	CLKCSR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x84)))
	CLKSELR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x85)))
	BGCCR   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x81)))
	BGCRR   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x80)))

	// EEPROM
	EEARL = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3e)))
	EEARH = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3f)))
	EEDR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3d)))
	EECR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3c)))

	// Power Stage Controller
	PICR0L  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x68)))
	PICR0H  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x69)))
	PFRC0B  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x63)))
	PFRC0A  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x62)))
	PCTL0   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x32)))
	PCNF0   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x31)))
	OCR0RBL = (*volatile.Register8)(unsafe.Pointer(uintptr(0x44)))
	OCR0RBH = (*volatile.Register8)(unsafe.Pointer(uintptr(0x45)))
	OCR0SBL = (*volatile.Register8)(unsafe.Pointer(uintptr(0x42)))
	OCR0SBH = (*volatile.Register8)(unsafe.Pointer(uintptr(0x43)))
	OCR0RAL = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4a)))
	OCR0RAH = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4b)))
	OCR0SAL = (*volatile.Register8)(unsafe.Pointer(uintptr(0x60)))
	OCR0SAH = (*volatile.Register8)(unsafe.Pointer(uintptr(0x61)))
	PSOC0   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x6a)))
	PIM0    = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2f)))
	PIFR0   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x30)))
	PICR2L  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x6c)))
	PICR2H  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x6d)))
	PFRC2B  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x67)))
	PFRC2A  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x66)))
	PCTL2   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x36)))
	PCNF2   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x35)))
	PCNFE2  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x70)))
	OCR2RBL = (*volatile.Register8)(unsafe.Pointer(uintptr(0x48)))
	OCR2RBH = (*volatile.Register8)(unsafe.Pointer(uintptr(0x49)))
	OCR2SBL = (*volatile.Register8)(unsafe.Pointer(uintptr(0x46)))
	OCR2SBH = (*volatile.Register8)(unsafe.Pointer(uintptr(0x47)))
	OCR2RAL = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4e)))
	OCR2RAH = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4f)))
	OCR2SAL = (*volatile.Register8)(unsafe.Pointer(uintptr(0x64)))
	OCR2SAH = (*volatile.Register8)(unsafe.Pointer(uintptr(0x65)))
	POM2    = (*volatile.Register8)(unsafe.Pointer(uintptr(0x6f)))
	PSOC2   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x6e)))
	PIM2    = (*volatile.Register8)(unsafe.Pointer(uintptr(0x33)))
	PIFR2   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x34)))
	PASDLY2 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x71)))

	// Timer/Counter, 16-bit
	TIMSK1 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x21)))
	TIFR1  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x22)))
	TCCR1B = (*volatile.Register8)(unsafe.Pointer(uintptr(0x8a)))
	TCNT1L = (*volatile.Register8)(unsafe.Pointer(uintptr(0x5a)))
	TCNT1H = (*volatile.Register8)(unsafe.Pointer(uintptr(0x5b)))
	ICR1L  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x8c)))
	ICR1H  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x8d)))

	// Bootloader
	SPMCSR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x57)))
)

// Bitfields for FUSE: Fuses
const (
	// EXTENDED
	EXTENDED_PSC2RB    = 0x80 // PSC2 Reset Behavior
	EXTENDED_PSC2RBA   = 0x40 // PSC2 Reset Behavior for 22 and 23
	EXTENDED_PSC0RB    = 0x20 // PSC0 Reset Behavior
	EXTENDED_PSCRV     = 0x10 // PSC Reset Value
	EXTENDED_PSCINRB   = 0x8  // PSC2 and PSC0 input Reset Behavior
	EXTENDED_BODLEVEL0 = 0x1  // Brown-out Detector Trigger Level
	EXTENDED_BODLEVEL1 = 0x2  // Brown-out Detector Trigger Level
	EXTENDED_BODLEVEL2 = 0x4  // Brown-out Detector Trigger Level

	// HIGH
	HIGH_RSTDISBL = 0x80 // Reset Disabled (Enable PE0 as I/O pin)
	HIGH_DWEN     = 0x40 // Debug Wire enable
	HIGH_SPIEN    = 0x20 // Serial program downloading (SPI) enabled
	HIGH_WDTON    = 0x10 // Watch-dog Timer always on
	HIGH_EESAVE   = 0x8  // Preserve EEPROM through the Chip Erase cycle
	HIGH_BOOTSZ0  = 0x2  // Select Boot Size
	HIGH_BOOTSZ1  = 0x4  // Select Boot Size
	HIGH_BOOTRST  = 0x1  // Select Reset Vector

	// LOW
	LOW_CKDIV8     = 0x80 // Divide clock by 8 internally
	LOW_CKOUT      = 0x40 // Clock output on PORTD1
	LOW_SUT_CKSEL0 = 0x1  // Select Clock Source
	LOW_SUT_CKSEL1 = 0x2  // Select Clock Source
	LOW_SUT_CKSEL2 = 0x4  // Select Clock Source
	LOW_SUT_CKSEL3 = 0x8  // Select Clock Source
	LOW_SUT_CKSEL4 = 0x10 // Select Clock Source
	LOW_SUT_CKSEL5 = 0x20 // Select Clock Source
)

// Bitfields for LOCKBIT: Lockbits
const (
	// LOCKBIT
	LOCKBIT_LB0   = 0x1  // Memory Lock
	LOCKBIT_LB1   = 0x2  // Memory Lock
	LOCKBIT_BLB00 = 0x4  // Boot Loader Protection Mode
	LOCKBIT_BLB01 = 0x8  // Boot Loader Protection Mode
	LOCKBIT_BLB10 = 0x10 // Boot Loader Protection Mode
	LOCKBIT_BLB11 = 0x20 // Boot Loader Protection Mode
)

// Bitfields for DAC: Digital-to-Analog Converter
const (
	// DACL: DAC Data Register

	// DACH: DAC Data Register
	DAC_DACH0 = 0x1  // DAC Data Register Bits
	DAC_DACH1 = 0x2  // DAC Data Register Bits
	DAC_DACH2 = 0x4  // DAC Data Register Bits
	DAC_DACH3 = 0x8  // DAC Data Register Bits
	DAC_DACH4 = 0x10 // DAC Data Register Bits
	DAC_DACH5 = 0x20 // DAC Data Register Bits
	DAC_DACH6 = 0x40 // DAC Data Register Bits
	DAC_DACH7 = 0x80 // DAC Data Register Bits

	// DACON: DAC Control Register
	DACON_DAATE = 0x80 // DAC Auto Trigger Enable Bit
	DACON_DATS0 = 0x10 // DAC Trigger Selection Bits
	DACON_DATS1 = 0x20 // DAC Trigger Selection Bits
	DACON_DATS2 = 0x40 // DAC Trigger Selection Bits
	DACON_DALA  = 0x4  // DAC Left Adjust
	DACON_DAEN  = 0x1  // DAC Enable Bit
)

// Bitfields for SPI: Serial Peripheral Interface
const (
	// SPCR: SPI Control Register
	SPCR_SPIE = 0x80 // SPI Interrupt Enable
	SPCR_SPE  = 0x40 // SPI Enable
	SPCR_DORD = 0x20 // Data Order
	SPCR_MSTR = 0x10 // Master/Slave Select
	SPCR_CPOL = 0x8  // Clock polarity
	SPCR_CPHA = 0x4  // Clock Phase
	SPCR_SPR0 = 0x1  // SPI Clock Rate Selects
	SPCR_SPR1 = 0x2  // SPI Clock Rate Selects

	// SPSR: SPI Status Register
	SPSR_SPIF  = 0x80 // SPI Interrupt Flag
	SPSR_WCOL  = 0x40 // Write Collision Flag
	SPSR_SPI2X = 0x1  // Double SPI Speed Bit

	// SPDR: SPI Data Register
	SPDR_SPD0 = 0x1  // SPI Data bits
	SPDR_SPD1 = 0x2  // SPI Data bits
	SPDR_SPD2 = 0x4  // SPI Data bits
	SPDR_SPD3 = 0x8  // SPI Data bits
	SPDR_SPD4 = 0x10 // SPI Data bits
	SPDR_SPD5 = 0x20 // SPI Data bits
	SPDR_SPD6 = 0x40 // SPI Data bits
	SPDR_SPD7 = 0x80 // SPI Data bits
)

// Bitfields for WDT: Watchdog Timer
const (
	// WDTCSR: Watchdog Timer Control Register
	WDTCSR_WDIF = 0x80 // Watchdog Timeout Interrupt Flag
	WDTCSR_WDIE = 0x40 // Watchdog Timeout Interrupt Enable
	WDTCSR_WDP0 = 0x1  // Watchdog Timer Prescaler Bits
	WDTCSR_WDP1 = 0x2  // Watchdog Timer Prescaler Bits
	WDTCSR_WDP2 = 0x4  // Watchdog Timer Prescaler Bits
	WDTCSR_WDP3 = 0x20 // Watchdog Timer Prescaler Bits
	WDTCSR_WDCE = 0x10 // Watchdog Change Enable
	WDTCSR_WDE  = 0x8  // Watch Dog Enable
)

// Bitfields for EXINT: External Interrupts
const (
	// EICRA: External Interrupt Control Register A
	EICRA_ISC20 = 0x10 // External Interrupt Sense Control Bit
	EICRA_ISC21 = 0x20 // External Interrupt Sense Control Bit
	EICRA_ISC10 = 0x4  // External Interrupt Sense Control Bit
	EICRA_ISC11 = 0x8  // External Interrupt Sense Control Bit
	EICRA_ISC00 = 0x1  // External Interrupt Sense Control Bit
	EICRA_ISC01 = 0x2  // External Interrupt Sense Control Bit

	// EIMSK: External Interrupt Mask Register
	EIMSK_INT0 = 0x1 // External Interrupt Request 2 Enable
	EIMSK_INT1 = 0x2 // External Interrupt Request 2 Enable
	EIMSK_INT2 = 0x4 // External Interrupt Request 2 Enable

	// EIFR: External Interrupt Flag Register
	EIFR_INTF0 = 0x1 // External Interrupt Flags
	EIFR_INTF1 = 0x2 // External Interrupt Flags
	EIFR_INTF2 = 0x4 // External Interrupt Flags
)

// Bitfields for ADC: Analog-to-Digital Converter
const (
	// ADMUX: The ADC multiplexer Selection Register
	ADMUX_REFS0 = 0x40 // Reference Selection Bits
	ADMUX_REFS1 = 0x80 // Reference Selection Bits
	ADMUX_ADLAR = 0x20 // Left Adjust Result
	ADMUX_MUX0  = 0x1  // Analog Channel and Gain Selection Bits
	ADMUX_MUX1  = 0x2  // Analog Channel and Gain Selection Bits
	ADMUX_MUX2  = 0x4  // Analog Channel and Gain Selection Bits
	ADMUX_MUX3  = 0x8  // Analog Channel and Gain Selection Bits

	// ADCSRA: The ADC Control and Status register
	ADCSRA_ADEN  = 0x80 // ADC Enable
	ADCSRA_ADSC  = 0x40 // ADC Start Conversion
	ADCSRA_ADATE = 0x20 // ADC Auto Trigger Enable
	ADCSRA_ADIF  = 0x10 // ADC Interrupt Flag
	ADCSRA_ADIE  = 0x8  // ADC Interrupt Enable
	ADCSRA_ADPS0 = 0x1  // ADC Prescaler Select Bits
	ADCSRA_ADPS1 = 0x2  // ADC Prescaler Select Bits
	ADCSRA_ADPS2 = 0x4  // ADC Prescaler Select Bits

	// ADCL: ADC Data Register Bytes

	// ADCH: ADC Data Register Bytes
	ADC_ADC0 = 0x1  // ADC Data Register
	ADC_ADC1 = 0x2  // ADC Data Register
	ADC_ADC2 = 0x4  // ADC Data Register
	ADC_ADC3 = 0x8  // ADC Data Register
	ADC_ADC4 = 0x10 // ADC Data Register
	ADC_ADC5 = 0x20 // ADC Data Register
	ADC_ADC6 = 0x40 // ADC Data Register
	ADC_ADC7 = 0x80 // ADC Data Register

	// ADCSRB: ADC Control and Status Register B
	ADCSRB_ADHSM   = 0x80 // ADC High Speed Mode
	ADCSRB_ADNCDIS = 0x40 // ADC Noise Canceller Disable
	ADCSRB_ADSSEN  = 0x10 // ADC Single Shot Enable on PSC's Synchronisation Signals
	ADCSRB_ADTS0   = 0x1  // ADC Auto Trigger Sources
	ADCSRB_ADTS1   = 0x2  // ADC Auto Trigger Sources
	ADCSRB_ADTS2   = 0x4  // ADC Auto Trigger Sources
	ADCSRB_ADTS3   = 0x8  // ADC Auto Trigger Sources

	// DIDR0: Digital Input Disable Register 0
	DIDR0_ADC8D = 0x80 // ADC8 Digital input Disable
	DIDR0_ADC7D = 0x40 // ADC7 Digital input Disable
	DIDR0_ADC5D = 0x20 // ADC5 Digital input Disable
	DIDR0_ADC4D = 0x10 // ADC4 Digital input Disable
	DIDR0_ADC3D = 0x8  // ADC3 Digital input Disable
	DIDR0_ADC2D = 0x4  // ADC2 Digital input Disable
	DIDR0_ADC1D = 0x2  // ADC1 Digital input Disable
	DIDR0_ADC0D = 0x1  // ADC0 Digital input Disable

	// DIDR1: Digital Input Disable Register 1
	DIDR1_ACMP1MD = 0x8
	DIDR1_AMP0PD  = 0x4
	DIDR1_ADC10D  = 0x2 // ADC10 Digital input Disable
	DIDR1_ADC9D   = 0x1 // ADC9 Digital input Disable

	// AMP0CSR
	AMP0CSR_AMP0EN  = 0x80
	AMP0CSR_AMP0IS  = 0x40
	AMP0CSR_AMP0G0  = 0x10
	AMP0CSR_AMP0G1  = 0x20
	AMP0CSR_AMP0GS  = 0x8
	AMP0CSR_AMP0TS0 = 0x1
	AMP0CSR_AMP0TS1 = 0x2
)

// Bitfields for AC: Analog Comparator
const (
	// AC3CON: Analog Comparator3 Control Register
	AC3CON_AC3EN  = 0x80 // Analog Comparator3 Enable Bit
	AC3CON_AC3IE  = 0x40 // Analog Comparator 3 Interrupt Enable Bit
	AC3CON_AC3IS0 = 0x10 // Analog Comparator 3  Interrupt Select Bit
	AC3CON_AC3IS1 = 0x20 // Analog Comparator 3  Interrupt Select Bit
	AC3CON_AC3OEA = 0x8  // Analog Comparator 3 Alternate Output Enable
	AC3CON_AC3M0  = 0x1  // Analog Comparator 3 Multiplexer Register
	AC3CON_AC3M1  = 0x2  // Analog Comparator 3 Multiplexer Register
	AC3CON_AC3M2  = 0x4  // Analog Comparator 3 Multiplexer Register

	// AC1CON: Analog Comparator 1 Control Register
	AC1CON_AC1EN  = 0x80 // Analog Comparator 1 Enable Bit
	AC1CON_AC1IE  = 0x40 // Analog Comparator 1 Interrupt Enable Bit
	AC1CON_AC1IS0 = 0x10 // Analog Comparator 1  Interrupt Select Bit
	AC1CON_AC1IS1 = 0x20 // Analog Comparator 1  Interrupt Select Bit
	AC1CON_AC1M0  = 0x1  // Analog Comparator 1 Multiplexer Register
	AC1CON_AC1M1  = 0x2  // Analog Comparator 1 Multiplexer Register
	AC1CON_AC1M2  = 0x4  // Analog Comparator 1 Multiplexer Register

	// AC2CON: Analog Comparator 2 Control Register
	AC2CON_AC2EN  = 0x80 // Analog Comparator 2 Enable Bit
	AC2CON_AC2IE  = 0x40 // Analog Comparator 2 Interrupt Enable Bit
	AC2CON_AC2IS0 = 0x10 // Analog Comparator 2  Interrupt Select Bit
	AC2CON_AC2IS1 = 0x20 // Analog Comparator 2  Interrupt Select Bit
	AC2CON_AC2M0  = 0x1  // Analog Comparator 2 Multiplexer Register
	AC2CON_AC2M1  = 0x2  // Analog Comparator 2 Multiplexer Register
	AC2CON_AC2M2  = 0x4  // Analog Comparator 2 Multiplexer Register

	// ACSR: Analog Comparator Status Register
	ACSR_AC3IF = 0x80 // Analog Comparator 3 Interrupt Flag Bit
	ACSR_AC2IF = 0x40 // Analog Comparator 2 Interrupt Flag Bit
	ACSR_AC1IF = 0x20 // Analog Comparator 1  Interrupt Flag Bit
	ACSR_AC3O  = 0x8  // Analog Comparator 3 Output Bit
	ACSR_AC2O  = 0x4  // Analog Comparator 2 Output Bit
	ACSR_AC1O  = 0x2  // Analog Comparator 1 Output Bit

	// AC3ECON
	AC3ECON_AC3OI = 0x20 // Analog Comparator Ouput Invert
	AC3ECON_AC3OE = 0x10 // Analog Comparator Ouput Enable
	AC3ECON_AC3H0 = 0x1  // Analog Comparator Hysteresis Select
	AC3ECON_AC3H1 = 0x2  // Analog Comparator Hysteresis Select
	AC3ECON_AC3H2 = 0x4  // Analog Comparator Hysteresis Select

	// AC2ECON
	AC2ECON_AC2OI = 0x20 // Analog Comparator Ouput Invert
	AC2ECON_AC2OE = 0x10 // Analog Comparator Ouput Enable
	AC2ECON_AC2H0 = 0x1  // Analog Comparator Hysteresis Select
	AC2ECON_AC2H1 = 0x2  // Analog Comparator Hysteresis Select
	AC2ECON_AC2H2 = 0x4  // Analog Comparator Hysteresis Select

	// AC1ECON
	AC1ECON_AC1OI  = 0x20 // Analog Comparator Ouput Invert
	AC1ECON_AC1OE  = 0x10 // Analog Comparator Ouput Enable
	AC1ECON_AC1ICE = 0x8  // Analog Comparator Interrupt Capture Enable
	AC1ECON_AC1H0  = 0x1  // Analog Comparator Hysteresis Select
	AC1ECON_AC1H1  = 0x2  // Analog Comparator Hysteresis Select
	AC1ECON_AC1H2  = 0x4  // Analog Comparator Hysteresis Select
)

// Bitfields for CPU: CPU Registers
const (
	// SREG: Status Register
	SREG_I = 0x80 // Global Interrupt Enable
	SREG_T = 0x40 // Bit Copy Storage
	SREG_H = 0x20 // Half Carry Flag
	SREG_S = 0x10 // Sign Bit
	SREG_V = 0x8  // Two's Complement Overflow Flag
	SREG_N = 0x4  // Negative Flag
	SREG_Z = 0x2  // Zero Flag
	SREG_C = 0x1  // Carry Flag

	// MCUCR: MCU Control Register
	MCUCR_PUD    = 0x10 // Pull-up disable
	MCUCR_RSTDIS = 0x8  // Reset Pin Disable
	MCUCR_CKRC81 = 0x4  // Frequency Selection of the Calibrated RC Oscillator
	MCUCR_IVSEL  = 0x2  // Interrupt Vector Select
	MCUCR_IVCE   = 0x1  // Interrupt Vector Change Enable

	// MCUSR: MCU Status Register
	MCUSR_WDRF  = 0x8 // Watchdog Reset Flag
	MCUSR_BORF  = 0x4 // Brown-out Reset Flag
	MCUSR_EXTRF = 0x2 // External Reset Flag
	MCUSR_PORF  = 0x1 // Power-on reset flag

	// OSCCAL: Oscillator Calibration Value
	OSCCAL_OSCCAL0 = 0x1  // Oscillator Calibration
	OSCCAL_OSCCAL1 = 0x2  // Oscillator Calibration
	OSCCAL_OSCCAL2 = 0x4  // Oscillator Calibration
	OSCCAL_OSCCAL3 = 0x8  // Oscillator Calibration
	OSCCAL_OSCCAL4 = 0x10 // Oscillator Calibration
	OSCCAL_OSCCAL5 = 0x20 // Oscillator Calibration
	OSCCAL_OSCCAL6 = 0x40 // Oscillator Calibration
	OSCCAL_OSCCAL7 = 0x80 // Oscillator Calibration

	// CLKPR
	CLKPR_CLKPCE = 0x80
	CLKPR_CLKPS0 = 0x1
	CLKPR_CLKPS1 = 0x2
	CLKPR_CLKPS2 = 0x4
	CLKPR_CLKPS3 = 0x8

	// SMCR: Sleep Mode Control Register
	SMCR_SM0 = 0x2 // Sleep Mode Select bits
	SMCR_SM1 = 0x4 // Sleep Mode Select bits
	SMCR_SM2 = 0x8 // Sleep Mode Select bits
	SMCR_SE  = 0x1 // Sleep Enable

	// GPIOR2: General Purpose IO Register 2
	GPIOR2_GPIOR0 = 0x1  // General Purpose IO Register 2 bis
	GPIOR2_GPIOR1 = 0x2  // General Purpose IO Register 2 bis
	GPIOR2_GPIOR2 = 0x4  // General Purpose IO Register 2 bis
	GPIOR2_GPIOR3 = 0x8  // General Purpose IO Register 2 bis
	GPIOR2_GPIOR4 = 0x10 // General Purpose IO Register 2 bis
	GPIOR2_GPIOR5 = 0x20 // General Purpose IO Register 2 bis
	GPIOR2_GPIOR6 = 0x40 // General Purpose IO Register 2 bis
	GPIOR2_GPIOR7 = 0x80 // General Purpose IO Register 2 bis

	// GPIOR1: General Purpose IO Register 1
	GPIOR1_GPIOR0 = 0x1  // General Purpose IO Register 1 bis
	GPIOR1_GPIOR1 = 0x2  // General Purpose IO Register 1 bis
	GPIOR1_GPIOR2 = 0x4  // General Purpose IO Register 1 bis
	GPIOR1_GPIOR3 = 0x8  // General Purpose IO Register 1 bis
	GPIOR1_GPIOR4 = 0x10 // General Purpose IO Register 1 bis
	GPIOR1_GPIOR5 = 0x20 // General Purpose IO Register 1 bis
	GPIOR1_GPIOR6 = 0x40 // General Purpose IO Register 1 bis
	GPIOR1_GPIOR7 = 0x80 // General Purpose IO Register 1 bis

	// GPIOR0: General Purpose IO Register 0
	GPIOR0_GPIOR07 = 0x80 // General Purpose IO Register 0 bit 7
	GPIOR0_GPIOR06 = 0x40 // General Purpose IO Register 0 bit 6
	GPIOR0_GPIOR05 = 0x20 // General Purpose IO Register 0 bit 5
	GPIOR0_GPIOR04 = 0x10 // General Purpose IO Register 0 bit 4
	GPIOR0_GPIOR03 = 0x8  // General Purpose IO Register 0 bit 3
	GPIOR0_GPIOR02 = 0x4  // General Purpose IO Register 0 bit 2
	GPIOR0_GPIOR01 = 0x2  // General Purpose IO Register 0 bit 1
	GPIOR0_GPIOR00 = 0x1  // General Purpose IO Register 0 bit 0

	// PLLCSR: PLL Control And Status Register
	PLLCSR_PLLF0 = 0x4
	PLLCSR_PLLF1 = 0x8
	PLLCSR_PLLF2 = 0x10
	PLLCSR_PLLF3 = 0x20
	PLLCSR_PLLE  = 0x2 // PLL Enable
	PLLCSR_PLOCK = 0x1 // PLL Lock Detector

	// PRR: Power Reduction Register
	PRR_PRPSC2 = 0x80 // Power Reduction PSC2
	PRR_PRPSCR = 0x20 // Power Reduction PSC0
	PRR_PRTIM1 = 0x10 // Power Reduction Timer/Counter1
	PRR_PRSPI  = 0x4  // Power Reduction Serial Peripheral Interface
	PRR_PRADC  = 0x1  // Power Reduction ADC

	// CLKCSR
	CLKCSR_CLKCCE = 0x80 // Clock Control Change Enable
	CLKCSR_CLKRDY = 0x10 // Clock Ready Flag
	CLKCSR_CLKC0  = 0x1  // Clock Control
	CLKCSR_CLKC1  = 0x2  // Clock Control
	CLKCSR_CLKC2  = 0x4  // Clock Control
	CLKCSR_CLKC3  = 0x8  // Clock Control

	// CLKSELR
	CLKSELR_COUT   = 0x40 // Clock OUT
	CLKSELR_CSUT0  = 0x10 // Clock Start up Time
	CLKSELR_CSUT1  = 0x20 // Clock Start up Time
	CLKSELR_CKSEL0 = 0x1  // Clock Source Select
	CLKSELR_CKSEL1 = 0x2  // Clock Source Select
	CLKSELR_CKSEL2 = 0x4  // Clock Source Select
	CLKSELR_CKSEL3 = 0x8  // Clock Source Select

	// BGCCR: BandGap Current Calibration Register
	BGCCR_BGCC0 = 0x1
	BGCCR_BGCC1 = 0x2
	BGCCR_BGCC2 = 0x4
	BGCCR_BGCC3 = 0x8

	// BGCRR: BandGap Resistor Calibration Register
	BGCRR_BGCR0 = 0x1
	BGCRR_BGCR1 = 0x2
	BGCRR_BGCR2 = 0x4
	BGCRR_BGCR3 = 0x8
)

// Bitfields for EEPROM: EEPROM
const (
	// EEARL: EEPROM Read/Write Access Bytes

	// EEARH: EEPROM Read/Write Access Bytes
	EEAR_EEAR0 = 0x1  // EEPROM Address bytes
	EEAR_EEAR1 = 0x2  // EEPROM Address bytes
	EEAR_EEAR2 = 0x4  // EEPROM Address bytes
	EEAR_EEAR3 = 0x8  // EEPROM Address bytes
	EEAR_EEAR4 = 0x10 // EEPROM Address bytes
	EEAR_EEAR5 = 0x20 // EEPROM Address bytes
	EEAR_EEAR6 = 0x40 // EEPROM Address bytes
	EEAR_EEAR7 = 0x80 // EEPROM Address bytes

	// EEDR: EEPROM Data Register
	EEDR_EEDR0 = 0x1  // EEPROM Data bits
	EEDR_EEDR1 = 0x2  // EEPROM Data bits
	EEDR_EEDR2 = 0x4  // EEPROM Data bits
	EEDR_EEDR3 = 0x8  // EEPROM Data bits
	EEDR_EEDR4 = 0x10 // EEPROM Data bits
	EEDR_EEDR5 = 0x20 // EEPROM Data bits
	EEDR_EEDR6 = 0x40 // EEPROM Data bits
	EEDR_EEDR7 = 0x80 // EEPROM Data bits

	// EECR: EEPROM Control Register
	EECR_NVMBSY = 0x80 // None Volatile Busy Memory Busy
	EECR_EEPAGE = 0x40 // EEPROM Page Access
	EECR_EEPM0  = 0x10 // EEPROM Programming Mode
	EECR_EEPM1  = 0x20 // EEPROM Programming Mode
	EECR_EERIE  = 0x8  // EEPROM Ready Interrupt Enable
	EECR_EEMWE  = 0x4  // EEPROM Master Write Enable
	EECR_EEWE   = 0x2  // EEPROM Write Enable
	EECR_EERE   = 0x1  // EEPROM Read Enable
)

// Bitfields for PSC: Power Stage Controller
const (
	// PICR0L: PSC 0 Input Capture Register

	// PICR0H: PSC 0 Input Capture Register
	PICR0_PCST0  = 0x8000 // PSC 0 Capture Software Trigger Bit
	PICR0_PICR00 = 0x1    // PSC 0 Input Capture Bytes
	PICR0_PICR01 = 0x2    // PSC 0 Input Capture Bytes
	PICR0_PICR02 = 0x4    // PSC 0 Input Capture Bytes
	PICR0_PICR03 = 0x8    // PSC 0 Input Capture Bytes
	PICR0_PICR04 = 0x10   // PSC 0 Input Capture Bytes
	PICR0_PICR05 = 0x20   // PSC 0 Input Capture Bytes
	PICR0_PICR06 = 0x40   // PSC 0 Input Capture Bytes
	PICR0_PICR07 = 0x80   // PSC 0 Input Capture Bytes

	// PFRC0B: PSC 0 Input B Control
	PFRC0B_PCAE0B  = 0x80 // PSC 0 Capture Enable Input Part B
	PFRC0B_PISEL0B = 0x40 // PSC 0 Input Select for Part B
	PFRC0B_PELEV0B = 0x20 // PSC 0 Edge Level Selector on Input Part B
	PFRC0B_PFLTE0B = 0x10 // PSC 0 Filter Enable on Input Part B
	PFRC0B_PRFM0B0 = 0x1  // PSC 0 Retrigger and Fault Mode for Part B
	PFRC0B_PRFM0B1 = 0x2  // PSC 0 Retrigger and Fault Mode for Part B
	PFRC0B_PRFM0B2 = 0x4  // PSC 0 Retrigger and Fault Mode for Part B
	PFRC0B_PRFM0B3 = 0x8  // PSC 0 Retrigger and Fault Mode for Part B

	// PFRC0A: PSC 0 Input A Control
	PFRC0A_PCAE0A  = 0x80 // PSC 0 Capture Enable Input Part A
	PFRC0A_PISEL0A = 0x40 // PSC 0 Input Select for Part A
	PFRC0A_PELEV0A = 0x20 // PSC 0 Edge Level Selector on Input Part A
	PFRC0A_PFLTE0A = 0x10 // PSC 0 Filter Enable on Input Part A
	PFRC0A_PRFM0A0 = 0x1  // PSC 0 Retrigger and Fault Mode for Part A
	PFRC0A_PRFM0A1 = 0x2  // PSC 0 Retrigger and Fault Mode for Part A
	PFRC0A_PRFM0A2 = 0x4  // PSC 0 Retrigger and Fault Mode for Part A
	PFRC0A_PRFM0A3 = 0x8  // PSC 0 Retrigger and Fault Mode for Part A

	// PCTL0: PSC 0 Control Register
	PCTL0_PPRE00 = 0x40 // PSC 0 Prescaler Selects
	PCTL0_PPRE01 = 0x80 // PSC 0 Prescaler Selects
	PCTL0_PBFM00 = 0x4  // PSC 0 Balance Flank Width Modulation
	PCTL0_PBFM01 = 0x20 // PSC 0 Balance Flank Width Modulation
	PCTL0_PAOC0B = 0x10 // PSC 0 Asynchronous Output Control B
	PCTL0_PAOC0A = 0x8  // PSC 0 Asynchronous Output Control A
	PCTL0_PCCYC0 = 0x2  // PSC0 Complete Cycle
	PCTL0_PRUN0  = 0x1  // PSC 0 Run

	// PCNF0: PSC 0 Configuration Register
	PCNF0_PFIFTY0  = 0x80 // PSC 0 Fifty
	PCNF0_PALOCK0  = 0x40 // PSC 0 Autolock
	PCNF0_PLOCK0   = 0x20 // PSC 0 Lock
	PCNF0_PMODE00  = 0x8  // PSC 0 Mode
	PCNF0_PMODE01  = 0x10 // PSC 0 Mode
	PCNF0_POP0     = 0x4  // PSC 0 Output Polarity
	PCNF0_PCLKSEL0 = 0x2  // PSC 0 Input Clock Select

	// OCR0RBL: Output Compare RB Register

	// OCR0RBH: Output Compare RB Register
	OCR0RB_OCR0RB0 = 0x1  // Output Compare 0 RB
	OCR0RB_OCR0RB1 = 0x2  // Output Compare 0 RB
	OCR0RB_OCR0RB2 = 0x4  // Output Compare 0 RB
	OCR0RB_OCR0RB3 = 0x8  // Output Compare 0 RB
	OCR0RB_OCR0RB4 = 0x10 // Output Compare 0 RB
	OCR0RB_OCR0RB5 = 0x20 // Output Compare 0 RB
	OCR0RB_OCR0RB6 = 0x40 // Output Compare 0 RB
	OCR0RB_OCR0RB7 = 0x80 // Output Compare 0 RB

	// OCR0SBL: Output Compare SB Register

	// OCR0SBH: Output Compare SB Register
	OCR0SB_OCR0SB0 = 0x1  // Output Compare 0 SB
	OCR0SB_OCR0SB1 = 0x2  // Output Compare 0 SB
	OCR0SB_OCR0SB2 = 0x4  // Output Compare 0 SB
	OCR0SB_OCR0SB3 = 0x8  // Output Compare 0 SB
	OCR0SB_OCR0SB4 = 0x10 // Output Compare 0 SB
	OCR0SB_OCR0SB5 = 0x20 // Output Compare 0 SB
	OCR0SB_OCR0SB6 = 0x40 // Output Compare 0 SB
	OCR0SB_OCR0SB7 = 0x80 // Output Compare 0 SB

	// OCR0RAL: Output Compare RA Register

	// OCR0RAH: Output Compare RA Register
	OCR0RA_OCR0RA0 = 0x1  // Output Compare 0 RA
	OCR0RA_OCR0RA1 = 0x2  // Output Compare 0 RA
	OCR0RA_OCR0RA2 = 0x4  // Output Compare 0 RA
	OCR0RA_OCR0RA3 = 0x8  // Output Compare 0 RA
	OCR0RA_OCR0RA4 = 0x10 // Output Compare 0 RA
	OCR0RA_OCR0RA5 = 0x20 // Output Compare 0 RA
	OCR0RA_OCR0RA6 = 0x40 // Output Compare 0 RA
	OCR0RA_OCR0RA7 = 0x80 // Output Compare 0 RA

	// OCR0SAL: Output Compare SA Register

	// OCR0SAH: Output Compare SA Register
	OCR0SA_OCR0SA0 = 0x1  // Output Compare 0 SA
	OCR0SA_OCR0SA1 = 0x2  // Output Compare 0 SA
	OCR0SA_OCR0SA2 = 0x4  // Output Compare 0 SA
	OCR0SA_OCR0SA3 = 0x8  // Output Compare 0 SA
	OCR0SA_OCR0SA4 = 0x10 // Output Compare 0 SA
	OCR0SA_OCR0SA5 = 0x20 // Output Compare 0 SA
	OCR0SA_OCR0SA6 = 0x40 // Output Compare 0 SA
	OCR0SA_OCR0SA7 = 0x80 // Output Compare 0 SA

	// PSOC0: PSC0 Synchro and Output Configuration
	PSOC0_PISEL0A1 = 0x80 // PSC Input Select
	PSOC0_PISEL0B1 = 0x40 // PSC Input Select
	PSOC0_PSYNC00  = 0x10 // Synchronisation out for ADC selection
	PSOC0_PSYNC01  = 0x20 // Synchronisation out for ADC selection
	PSOC0_POEN0B   = 0x4  // PSCOUT01 Output Enable
	PSOC0_POEN0A   = 0x1  // PSCOUT00 Output Enable

	// PIM0: PSC0 Interrupt Mask Register
	PIM0_PEVE0B  = 0x10 // External Event B Interrupt Enable
	PIM0_PEVE0A  = 0x8  // External Event A Interrupt Enable
	PIM0_PEOEPE0 = 0x2  // End of Enhanced Cycle Enable
	PIM0_PEOPE0  = 0x1  // End of Cycle Interrupt Enable

	// PIFR0: PSC0 Interrupt Flag Register
	PIFR0_POAC0B = 0x80 // PSC 0 Output A Activity
	PIFR0_POAC0A = 0x40 // PSC 0 Output A Activity
	PIFR0_PEV0B  = 0x10 // External Event B Interrupt
	PIFR0_PEV0A  = 0x8  // External Event A Interrupt
	PIFR0_PRN00  = 0x2  // Ramp Number
	PIFR0_PRN01  = 0x4  // Ramp Number
	PIFR0_PEOP0  = 0x1  // End of PSC0 Interrupt

	// PICR2L: PSC 2 Input Capture Register

	// PICR2H: PSC 2 Input Capture Register
	PICR2_PCST2  = 0x8000 // PSC 2 Capture Software Trigger Bit
	PICR2_PICR20 = 0x1    // PSC 2 Input Capture Bytes
	PICR2_PICR21 = 0x2    // PSC 2 Input Capture Bytes
	PICR2_PICR22 = 0x4    // PSC 2 Input Capture Bytes
	PICR2_PICR23 = 0x8    // PSC 2 Input Capture Bytes
	PICR2_PICR24 = 0x10   // PSC 2 Input Capture Bytes
	PICR2_PICR25 = 0x20   // PSC 2 Input Capture Bytes
	PICR2_PICR26 = 0x40   // PSC 2 Input Capture Bytes
	PICR2_PICR27 = 0x80   // PSC 2 Input Capture Bytes

	// PFRC2B: PSC 2 Input B Control
	PFRC2B_PCAE2B  = 0x80 // PSC 2 Capture Enable Input Part B
	PFRC2B_PISEL2B = 0x40 // PSC 2 Input Select for Part B
	PFRC2B_PELEV2B = 0x20 // PSC 2 Edge Level Selector on Input Part B
	PFRC2B_PFLTE2B = 0x10 // PSC 2 Filter Enable on Input Part B
	PFRC2B_PRFM2B0 = 0x1  // PSC 2 Retrigger and Fault Mode for Part B
	PFRC2B_PRFM2B1 = 0x2  // PSC 2 Retrigger and Fault Mode for Part B
	PFRC2B_PRFM2B2 = 0x4  // PSC 2 Retrigger and Fault Mode for Part B
	PFRC2B_PRFM2B3 = 0x8  // PSC 2 Retrigger and Fault Mode for Part B

	// PFRC2A: PSC 2 Input B Control
	PFRC2A_PCAE2A  = 0x80 // PSC 2 Capture Enable Input Part A
	PFRC2A_PISEL2A = 0x40 // PSC 2 Input Select for Part A
	PFRC2A_PELEV2A = 0x20 // PSC 2 Edge Level Selector on Input Part A
	PFRC2A_PFLTE2A = 0x10 // PSC 2 Filter Enable on Input Part A
	PFRC2A_PRFM2A0 = 0x1  // PSC 2 Retrigger and Fault Mode for Part A
	PFRC2A_PRFM2A1 = 0x2  // PSC 2 Retrigger and Fault Mode for Part A
	PFRC2A_PRFM2A2 = 0x4  // PSC 2 Retrigger and Fault Mode for Part A
	PFRC2A_PRFM2A3 = 0x8  // PSC 2 Retrigger and Fault Mode for Part A

	// PCTL2: PSC 2 Control Register
	PCTL2_PPRE20 = 0x40 // PSC 2 Prescaler Selects
	PCTL2_PPRE21 = 0x80 // PSC 2 Prescaler Selects
	PCTL2_PBFM2  = 0x20 // Balance Flank Width Modulation
	PCTL2_PAOC2B = 0x10 // PSC 2 Asynchronous Output Control B
	PCTL2_PAOC2A = 0x8  // PSC 2 Asynchronous Output Control A
	PCTL2_PARUN2 = 0x4  // PSC2 Auto Run
	PCTL2_PCCYC2 = 0x2  // PSC2 Complete Cycle
	PCTL2_PRUN2  = 0x1  // PSC 2 Run

	// PCNF2: PSC 2 Configuration Register
	PCNF2_PFIFTY2  = 0x80 // PSC 2 Fifty
	PCNF2_PALOCK2  = 0x40 // PSC 2 Autolock
	PCNF2_PLOCK2   = 0x20 // PSC 2 Lock
	PCNF2_PMODE20  = 0x8  // PSC 2 Mode
	PCNF2_PMODE21  = 0x10 // PSC 2 Mode
	PCNF2_POP2     = 0x4  // PSC 2 Output Polarity
	PCNF2_PCLKSEL2 = 0x2  // PSC 2 Input Clock Select
	PCNF2_POME2    = 0x1  // PSC 2 Output Matrix Enable

	// PCNFE2: PSC 2 Enhanced Configuration Register
	PCNFE2_PASDLK20 = 0x20
	PCNFE2_PASDLK21 = 0x40
	PCNFE2_PASDLK22 = 0x80
	PCNFE2_PBFM21   = 0x10
	PCNFE2_PELEV2A1 = 0x8
	PCNFE2_PELEV2B1 = 0x4
	PCNFE2_PISEL2A1 = 0x2
	PCNFE2_PISEL2B1 = 0x1

	// OCR2RBL: Output Compare RB Register

	// OCR2RBH: Output Compare RB Register
	OCR2RB_OCR2RB0 = 0x1  // Output Compare 2 RB
	OCR2RB_OCR2RB1 = 0x2  // Output Compare 2 RB
	OCR2RB_OCR2RB2 = 0x4  // Output Compare 2 RB
	OCR2RB_OCR2RB3 = 0x8  // Output Compare 2 RB
	OCR2RB_OCR2RB4 = 0x10 // Output Compare 2 RB
	OCR2RB_OCR2RB5 = 0x20 // Output Compare 2 RB
	OCR2RB_OCR2RB6 = 0x40 // Output Compare 2 RB
	OCR2RB_OCR2RB7 = 0x80 // Output Compare 2 RB

	// OCR2SBL: Output Compare SB Register

	// OCR2SBH: Output Compare SB Register
	OCR2SB_OCR2SB0 = 0x1  // Output Compare 2 SB
	OCR2SB_OCR2SB1 = 0x2  // Output Compare 2 SB
	OCR2SB_OCR2SB2 = 0x4  // Output Compare 2 SB
	OCR2SB_OCR2SB3 = 0x8  // Output Compare 2 SB
	OCR2SB_OCR2SB4 = 0x10 // Output Compare 2 SB
	OCR2SB_OCR2SB5 = 0x20 // Output Compare 2 SB
	OCR2SB_OCR2SB6 = 0x40 // Output Compare 2 SB
	OCR2SB_OCR2SB7 = 0x80 // Output Compare 2 SB

	// OCR2RAL: Output Compare RA Register

	// OCR2RAH: Output Compare RA Register
	OCR2RA_OCR2RA0 = 0x1  // Output Compare 2 RA
	OCR2RA_OCR2RA1 = 0x2  // Output Compare 2 RA
	OCR2RA_OCR2RA2 = 0x4  // Output Compare 2 RA
	OCR2RA_OCR2RA3 = 0x8  // Output Compare 2 RA
	OCR2RA_OCR2RA4 = 0x10 // Output Compare 2 RA
	OCR2RA_OCR2RA5 = 0x20 // Output Compare 2 RA
	OCR2RA_OCR2RA6 = 0x40 // Output Compare 2 RA
	OCR2RA_OCR2RA7 = 0x80 // Output Compare 2 RA

	// OCR2SAL: Output Compare SA Register

	// OCR2SAH: Output Compare SA Register
	OCR2SA_OCR2SA0 = 0x1  // Output Compare 2 SA
	OCR2SA_OCR2SA1 = 0x2  // Output Compare 2 SA
	OCR2SA_OCR2SA2 = 0x4  // Output Compare 2 SA
	OCR2SA_OCR2SA3 = 0x8  // Output Compare 2 SA
	OCR2SA_OCR2SA4 = 0x10 // Output Compare 2 SA
	OCR2SA_OCR2SA5 = 0x20 // Output Compare 2 SA
	OCR2SA_OCR2SA6 = 0x40 // Output Compare 2 SA
	OCR2SA_OCR2SA7 = 0x80 // Output Compare 2 SA

	// POM2: PSC 2 Output Matrix
	POM2_POMV2B0 = 0x10 // Output Matrix Output B Ramps
	POM2_POMV2B1 = 0x20 // Output Matrix Output B Ramps
	POM2_POMV2B2 = 0x40 // Output Matrix Output B Ramps
	POM2_POMV2B3 = 0x80 // Output Matrix Output B Ramps
	POM2_POMV2A0 = 0x1  // Output Matrix Output A Ramps
	POM2_POMV2A1 = 0x2  // Output Matrix Output A Ramps
	POM2_POMV2A2 = 0x4  // Output Matrix Output A Ramps
	POM2_POMV2A3 = 0x8  // Output Matrix Output A Ramps

	// PSOC2: PSC2 Synchro and Output Configuration
	PSOC2_POS20   = 0x40 // PSC 2 Output 23 Select
	PSOC2_POS21   = 0x80 // PSC 2 Output 23 Select
	PSOC2_PSYNC20 = 0x10 // Synchronization Out for ADC Selection
	PSOC2_PSYNC21 = 0x20 // Synchronization Out for ADC Selection
	PSOC2_POEN2D  = 0x8  // PSCOUT23 Output Enable
	PSOC2_POEN2B  = 0x4  // PSCOUT21 Output Enable
	PSOC2_POEN2C  = 0x2  // PSCOUT22 Output Enable
	PSOC2_POEN2A  = 0x1  // PSCOUT20 Output Enable

	// PIM2: PSC2 Interrupt Mask Register
	PIM2_PSEIE2  = 0x20 // PSC 2 Synchro Error Interrupt Enable
	PIM2_PEVE2B  = 0x10 // External Event B Interrupt Enable
	PIM2_PEVE2A  = 0x8  // External Event A Interrupt Enable
	PIM2_PEOEPE2 = 0x2  // End of Enhanced Cycle Interrupt Enable
	PIM2_PEOPE2  = 0x1  // End of Cycle Interrupt Enable

	// PIFR2: PSC2 Interrupt Flag Register
	PIFR2_POAC2B = 0x80 // PSC 2 Output A Activity
	PIFR2_POAC2A = 0x40 // PSC 2 Output A Activity
	PIFR2_PSEI2  = 0x20 // PSC 2 Synchro Error Interrupt
	PIFR2_PEV2B  = 0x10 // External Event B Interrupt
	PIFR2_PEV2A  = 0x8  // External Event A Interrupt
	PIFR2_PRN20  = 0x2  // Ramp Number
	PIFR2_PRN21  = 0x4  // Ramp Number
	PIFR2_PEOP2  = 0x1  // End of PSC2 Interrupt

	// PASDLY2: Analog Synchronization Delay Register
	PASDLY2_PASDLY20 = 0x1  // Analog Synchronization Delay bits
	PASDLY2_PASDLY21 = 0x2  // Analog Synchronization Delay bits
	PASDLY2_PASDLY22 = 0x4  // Analog Synchronization Delay bits
	PASDLY2_PASDLY23 = 0x8  // Analog Synchronization Delay bits
	PASDLY2_PASDLY24 = 0x10 // Analog Synchronization Delay bits
	PASDLY2_PASDLY25 = 0x20 // Analog Synchronization Delay bits
	PASDLY2_PASDLY26 = 0x40 // Analog Synchronization Delay bits
	PASDLY2_PASDLY27 = 0x80 // Analog Synchronization Delay bits
)

// Bitfields for TC16: Timer/Counter, 16-bit
const (
	// TIMSK1: Timer/Counter Interrupt Mask Register
	TIMSK1_ICIE1 = 0x20 // Timer/Counter1 Input Capture Interrupt Enable
	TIMSK1_TOIE1 = 0x1  // Timer/Counter1 Overflow Interrupt Enable

	// TIFR1: Timer/Counter Interrupt Flag register
	TIFR1_ICF1 = 0x20 // Input Capture Flag 1
	TIFR1_TOV1 = 0x1  // Timer/Counter1 Overflow Flag

	// TCCR1B: Timer/Counter1 Control Register B
	TCCR1B_ICNC1 = 0x80 // Input Capture 1 Noise Canceler
	TCCR1B_ICES1 = 0x40 // Input Capture 1 Edge Select
	TCCR1B_WGM13 = 0x10 // Waveform Generation Mode
	TCCR1B_CS10  = 0x1  // Prescaler source of Timer/Counter 1
	TCCR1B_CS11  = 0x2  // Prescaler source of Timer/Counter 1
	TCCR1B_CS12  = 0x4  // Prescaler source of Timer/Counter 1

	// TCNT1L: Timer/Counter1 Bytes

	// TCNT1H: Timer/Counter1 Bytes
	TCNT1_TCNT10 = 0x1  // Timer/Counter 1 bits
	TCNT1_TCNT11 = 0x2  // Timer/Counter 1 bits
	TCNT1_TCNT12 = 0x4  // Timer/Counter 1 bits
	TCNT1_TCNT13 = 0x8  // Timer/Counter 1 bits
	TCNT1_TCNT14 = 0x10 // Timer/Counter 1 bits
	TCNT1_TCNT15 = 0x20 // Timer/Counter 1 bits
	TCNT1_TCNT16 = 0x40 // Timer/Counter 1 bits
	TCNT1_TCNT17 = 0x80 // Timer/Counter 1 bits

	// ICR1L: Timer/Counter1 Input Capture Register Bytes

	// ICR1H: Timer/Counter1 Input Capture Register Bytes
	ICR1_ICR10 = 0x1  // Timer/Counter1 Input Capture bits
	ICR1_ICR11 = 0x2  // Timer/Counter1 Input Capture bits
	ICR1_ICR12 = 0x4  // Timer/Counter1 Input Capture bits
	ICR1_ICR13 = 0x8  // Timer/Counter1 Input Capture bits
	ICR1_ICR14 = 0x10 // Timer/Counter1 Input Capture bits
	ICR1_ICR15 = 0x20 // Timer/Counter1 Input Capture bits
	ICR1_ICR16 = 0x40 // Timer/Counter1 Input Capture bits
	ICR1_ICR17 = 0x80 // Timer/Counter1 Input Capture bits
)

// Bitfields for BOOT_LOAD: Bootloader
const (
	// SPMCSR: Store Program Memory Control Register
	SPMCSR_SPMIE  = 0x80 // SPM Interrupt Enable
	SPMCSR_RWWSB  = 0x40 // Read While Write Section Busy
	SPMCSR_SIGRD  = 0x20 // Signature Row Read
	SPMCSR_RWWSRE = 0x10 // Read While Write section read enable
	SPMCSR_BLBSET = 0x8  // Boot Lock Bit Set
	SPMCSR_PGWRT  = 0x4  // Page Write
	SPMCSR_PGERS  = 0x2  // Page Erase
	SPMCSR_SPMEN  = 0x1  // Store Program Memory Enable
)
