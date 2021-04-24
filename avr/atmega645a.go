// Automatically generated file. DO NOT EDIT.
// Generated by gen-device-avr.go from ATmega645A.atdf, see http://packs.download.atmel.com/

// +build avr,atmega645a

// Device information for the ATmega645A.
package avr

import (
	"runtime/interrupt"
	"runtime/volatile"
	"unsafe"
)

// Some information about this device.
const (
	DEVICE = "ATmega645A"
	ARCH   = "AVR8"
	FAMILY = "megaAVR"
)

// Interrupts
const (
	IRQ_RESET        = 0  // External Pin,Power-on Reset,Brown-out Reset,Watchdog Reset,and JTAG AVR Reset. See Datasheet.
	IRQ_INT0         = 1  // External Interrupt Request 0
	IRQ_PCINT0       = 2  // Pin Change Interrupt Request 0
	IRQ_PCINT1       = 3  // Pin Change Interrupt Request 1
	IRQ_TIMER2_COMP  = 4  // Timer/Counter2 Compare Match
	IRQ_TIMER2_OVF   = 5  // Timer/Counter2 Overflow
	IRQ_TIMER1_CAPT  = 6  // Timer/Counter1 Capture Event
	IRQ_TIMER1_COMPA = 7  // Timer/Counter1 Compare Match A
	IRQ_TIMER1_COMPB = 8  // Timer/Counter Compare Match B
	IRQ_TIMER1_OVF   = 9  // Timer/Counter1 Overflow
	IRQ_TIMER0_COMP  = 10 // Timer/Counter0 Compare Match
	IRQ_TIMER0_OVF   = 11 // Timer/Counter0 Overflow
	IRQ_SPI_STC      = 12 // SPI Serial Transfer Complete
	IRQ_USART0_RX    = 13 // USART0, Rx Complete
	IRQ_USART0_UDRE  = 14 // USART0 Data register Empty
	IRQ_USART0_TX    = 15 // USART0, Tx Complete
	IRQ_USI_START    = 16 // USI Start Condition
	IRQ_USI_OVERFLOW = 17 // USI Overflow
	IRQ_ANALOG_COMP  = 18 // Analog Comparator
	IRQ_ADC          = 19 // ADC Conversion Complete
	IRQ_EE_READY     = 20 // EEPROM Ready
	IRQ_SPM_READY    = 21 // Store Program Memory Read
	IRQ_max          = 21 // Highest interrupt number on this device.
)

// Map interrupt numbers to function names.
// These aren't real calls, they're removed by the compiler.
var (
	_ = interrupt.Register(IRQ_RESET, "__vector_RESET")
	_ = interrupt.Register(IRQ_INT0, "__vector_INT0")
	_ = interrupt.Register(IRQ_PCINT0, "__vector_PCINT0")
	_ = interrupt.Register(IRQ_PCINT1, "__vector_PCINT1")
	_ = interrupt.Register(IRQ_TIMER2_COMP, "__vector_TIMER2_COMP")
	_ = interrupt.Register(IRQ_TIMER2_OVF, "__vector_TIMER2_OVF")
	_ = interrupt.Register(IRQ_TIMER1_CAPT, "__vector_TIMER1_CAPT")
	_ = interrupt.Register(IRQ_TIMER1_COMPA, "__vector_TIMER1_COMPA")
	_ = interrupt.Register(IRQ_TIMER1_COMPB, "__vector_TIMER1_COMPB")
	_ = interrupt.Register(IRQ_TIMER1_OVF, "__vector_TIMER1_OVF")
	_ = interrupt.Register(IRQ_TIMER0_COMP, "__vector_TIMER0_COMP")
	_ = interrupt.Register(IRQ_TIMER0_OVF, "__vector_TIMER0_OVF")
	_ = interrupt.Register(IRQ_SPI_STC, "__vector_SPI_STC")
	_ = interrupt.Register(IRQ_USART0_RX, "__vector_USART0_RX")
	_ = interrupt.Register(IRQ_USART0_UDRE, "__vector_USART0_UDRE")
	_ = interrupt.Register(IRQ_USART0_TX, "__vector_USART0_TX")
	_ = interrupt.Register(IRQ_USI_START, "__vector_USI_START")
	_ = interrupt.Register(IRQ_USI_OVERFLOW, "__vector_USI_OVERFLOW")
	_ = interrupt.Register(IRQ_ANALOG_COMP, "__vector_ANALOG_COMP")
	_ = interrupt.Register(IRQ_ADC, "__vector_ADC")
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

	// Analog-to-Digital Converter
	ADMUX  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x7c)))
	ADCSRA = (*volatile.Register8)(unsafe.Pointer(uintptr(0x7a)))
	ADCL   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x78)))
	ADCH   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x79)))
	DIDR0  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x7e)))

	// Analog Comparator
	ACSR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x50)))
	DIDR1 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x7f)))

	// Serial Peripheral Interface
	SPCR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4c)))
	SPSR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4d)))
	SPDR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4e)))

	// Universal Serial Interface
	USIDR = (*volatile.Register8)(unsafe.Pointer(uintptr(0xba)))
	USISR = (*volatile.Register8)(unsafe.Pointer(uintptr(0xb9)))
	USICR = (*volatile.Register8)(unsafe.Pointer(uintptr(0xb8)))

	// CPU Registers
	SREG   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x5f)))
	SPL    = (*volatile.Register8)(unsafe.Pointer(uintptr(0x5d)))
	SPH    = (*volatile.Register8)(unsafe.Pointer(uintptr(0x5e)))
	OSCCAL = (*volatile.Register8)(unsafe.Pointer(uintptr(0x66)))
	CLKPR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x61)))
	PRR    = (*volatile.Register8)(unsafe.Pointer(uintptr(0x64)))
	SMCR   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x53)))
	GPIOR2 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4b)))
	GPIOR1 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4a)))
	GPIOR0 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3e)))

	// JTAG Interface
	OCDR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x51)))

	// External Interrupts
	EICRA  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x69)))
	EIMSK  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3d)))
	EIFR   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3c)))
	PCMSK1 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x6c)))
	PCMSK0 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x6b)))

	// EEPROM
	EEARL = (*volatile.Register8)(unsafe.Pointer(uintptr(0x41)))
	EEARH = (*volatile.Register8)(unsafe.Pointer(uintptr(0x42)))
	EEDR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x40)))
	EECR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3f)))

	// I/O Port
	PORTA = (*volatile.Register8)(unsafe.Pointer(uintptr(0x22)))
	DDRA  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x21)))
	PINA  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x20)))
	PORTB = (*volatile.Register8)(unsafe.Pointer(uintptr(0x25)))
	DDRB  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x24)))
	PINB  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x23)))
	PORTC = (*volatile.Register8)(unsafe.Pointer(uintptr(0x28)))
	DDRC  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x27)))
	PINC  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x26)))
	PORTD = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2b)))
	DDRD  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2a)))
	PIND  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x29)))
	PORTE = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2e)))
	DDRE  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2d)))
	PINE  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2c)))
	PORTF = (*volatile.Register8)(unsafe.Pointer(uintptr(0x31)))
	DDRF  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x30)))
	PINF  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2f)))
	PORTG = (*volatile.Register8)(unsafe.Pointer(uintptr(0x34)))
	DDRG  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x33)))
	PING  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x32)))

	// Timer/Counter, 8-bit
	TCCR0A = (*volatile.Register8)(unsafe.Pointer(uintptr(0x44)))
	TCNT0  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x46)))
	OCR0A  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x47)))
	TIMSK0 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x6e)))
	TIFR0  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x35)))

	// Timer/Counter, 8-bit Async
	TCCR2A = (*volatile.Register8)(unsafe.Pointer(uintptr(0xb0)))
	TCNT2  = (*volatile.Register8)(unsafe.Pointer(uintptr(0xb2)))
	OCR2A  = (*volatile.Register8)(unsafe.Pointer(uintptr(0xb3)))
	TIMSK2 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x70)))
	TIFR2  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x37)))
	ASSR   = (*volatile.Register8)(unsafe.Pointer(uintptr(0xb6)))

	// Timer/Counter, 16-bit
	TCCR1A = (*volatile.Register8)(unsafe.Pointer(uintptr(0x80)))
	TCCR1B = (*volatile.Register8)(unsafe.Pointer(uintptr(0x81)))
	TCCR1C = (*volatile.Register8)(unsafe.Pointer(uintptr(0x82)))
	TCNT1L = (*volatile.Register8)(unsafe.Pointer(uintptr(0x84)))
	TCNT1H = (*volatile.Register8)(unsafe.Pointer(uintptr(0x85)))
	OCR1AL = (*volatile.Register8)(unsafe.Pointer(uintptr(0x88)))
	OCR1AH = (*volatile.Register8)(unsafe.Pointer(uintptr(0x89)))
	OCR1BL = (*volatile.Register8)(unsafe.Pointer(uintptr(0x8a)))
	OCR1BH = (*volatile.Register8)(unsafe.Pointer(uintptr(0x8b)))
	ICR1L  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x86)))
	ICR1H  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x87)))
	TIMSK1 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x6f)))
	TIFR1  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x36)))

	// Watchdog Timer
	WDTCR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x60)))

	// Bootloader
	SPMCSR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x57)))

	// USART
	UDR0   = (*volatile.Register8)(unsafe.Pointer(uintptr(0xc6)))
	UCSR0A = (*volatile.Register8)(unsafe.Pointer(uintptr(0xc0)))
	UCSR0B = (*volatile.Register8)(unsafe.Pointer(uintptr(0xc1)))
	UCSR0C = (*volatile.Register8)(unsafe.Pointer(uintptr(0xc2)))
	UBRR0L = (*volatile.Register8)(unsafe.Pointer(uintptr(0xc4)))
	UBRR0H = (*volatile.Register8)(unsafe.Pointer(uintptr(0xc5)))
)

// Bitfields for FUSE: Fuses
const (
	// EXTENDED
	EXTENDED_BODLEVEL0 = 0x2 // Brown-out Detector trigger level
	EXTENDED_BODLEVEL1 = 0x4 // Brown-out Detector trigger level
	EXTENDED_RSTDISBL  = 0x1 // External Reset Disable

	// HIGH
	HIGH_OCDEN   = 0x80 // On-Chip Debug Enabled
	HIGH_JTAGEN  = 0x40 // JTAG Interface Enabled
	HIGH_SPIEN   = 0x20 // Serial program downloading (SPI) enable
	HIGH_WDTON   = 0x10 // Watchdog timer always on
	HIGH_EESAVE  = 0x8  // Preserve EEPROM through the Chip Erase cycle
	HIGH_BOOTSZ0 = 0x2  // Select Boot Size
	HIGH_BOOTSZ1 = 0x4  // Select Boot Size
	HIGH_BOOTRST = 0x1  // Boot Reset vector Enabled

	// LOW
	LOW_CKDIV8     = 0x80 // Divide clock by 8 internally
	LOW_CKOUT      = 0x40 // Clock output on PORTE7
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
	ADMUX_MUX4  = 0x10 // Analog Channel and Gain Selection Bits

	// ADCSRA: The ADC Control and Status register
	ADCSRA_ADEN  = 0x80 // ADC Enable
	ADCSRA_ADSC  = 0x40 // ADC Start Conversion
	ADCSRA_ADATE = 0x20 // ADC Auto Trigger Enable
	ADCSRA_ADIF  = 0x10 // ADC Interrupt Flag
	ADCSRA_ADIE  = 0x8  // ADC Interrupt Enable
	ADCSRA_ADPS0 = 0x1  // ADC  Prescaler Select Bits
	ADCSRA_ADPS1 = 0x2  // ADC  Prescaler Select Bits
	ADCSRA_ADPS2 = 0x4  // ADC  Prescaler Select Bits

	// DIDR0: Digital Input Disable Register 0
	DIDR0_ADC7D = 0x80 // ADC7 Digital input Disable
	DIDR0_ADC6D = 0x40 // ADC6 Digital input Disable
	DIDR0_ADC5D = 0x20 // ADC5 Digital input Disable
	DIDR0_ADC4D = 0x10 // ADC4 Digital input Disable
	DIDR0_ADC3D = 0x8  // ADC3 Digital input Disable
	DIDR0_ADC2D = 0x4  // ADC2 Digital input Disable
	DIDR0_ADC1D = 0x2  // ADC1 Digital input Disable
	DIDR0_ADC0D = 0x1  // ADC0 Digital input Disable
)

// Bitfields for AC: Analog Comparator
const (
	// ACSR: Analog Comparator Control And Status Register
	ACSR_ACD   = 0x80 // Analog Comparator Disable
	ACSR_ACBG  = 0x40 // Analog Comparator Bandgap Select
	ACSR_ACO   = 0x20 // Analog Compare Output
	ACSR_ACI   = 0x10 // Analog Comparator Interrupt Flag
	ACSR_ACIE  = 0x8  // Analog Comparator Interrupt Enable
	ACSR_ACIC  = 0x4  // Analog Comparator Input Capture Enable
	ACSR_ACIS0 = 0x1  // Analog Comparator Interrupt Mode Select bits
	ACSR_ACIS1 = 0x2  // Analog Comparator Interrupt Mode Select bits

	// DIDR1: Digital Input Disable Register 1
	DIDR1_AIN1D = 0x2 // AIN1 Digital Input Disable
	DIDR1_AIN0D = 0x1 // AIN0 Digital Input Disable
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
)

// Bitfields for USI: Universal Serial Interface
const (
	// USISR: USI Status Register
	USISR_USISIF  = 0x80 // Start Condition Interrupt Flag
	USISR_USIOIF  = 0x40 // Counter Overflow Interrupt Flag
	USISR_USIPF   = 0x20 // Stop Condition Flag
	USISR_USIDC   = 0x10 // Data Output Collision
	USISR_USICNT0 = 0x1  // USI Counter Value Bits
	USISR_USICNT1 = 0x2  // USI Counter Value Bits
	USISR_USICNT2 = 0x4  // USI Counter Value Bits
	USISR_USICNT3 = 0x8  // USI Counter Value Bits

	// USICR: USI Control Register
	USICR_USISIE = 0x80 // Start Condition Interrupt Enable
	USICR_USIOIE = 0x40 // Counter Overflow Interrupt Enable
	USICR_USIWM0 = 0x10 // USI Wire Mode Bits
	USICR_USIWM1 = 0x20 // USI Wire Mode Bits
	USICR_USICS0 = 0x4  // USI Clock Source Select Bits
	USICR_USICS1 = 0x8  // USI Clock Source Select Bits
	USICR_USICLK = 0x2  // Clock Strobe
	USICR_USITC  = 0x1  // Toggle Clock Port Pin
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

	// OSCCAL: Oscillator Calibration Value
	OSCCAL_OSCCAL0 = 0x1  // Oscillator Calibration
	OSCCAL_OSCCAL1 = 0x2  // Oscillator Calibration
	OSCCAL_OSCCAL2 = 0x4  // Oscillator Calibration
	OSCCAL_OSCCAL3 = 0x8  // Oscillator Calibration
	OSCCAL_OSCCAL4 = 0x10 // Oscillator Calibration
	OSCCAL_OSCCAL5 = 0x20 // Oscillator Calibration
	OSCCAL_OSCCAL6 = 0x40 // Oscillator Calibration
	OSCCAL_OSCCAL7 = 0x80 // Oscillator Calibration

	// CLKPR: Clock Prescale Register
	CLKPR_CLKPCE = 0x80 // Clock Prescaler Change Enable
	CLKPR_CLKPS0 = 0x1  // Clock Prescaler Select Bits
	CLKPR_CLKPS1 = 0x2  // Clock Prescaler Select Bits
	CLKPR_CLKPS2 = 0x4  // Clock Prescaler Select Bits
	CLKPR_CLKPS3 = 0x8  // Clock Prescaler Select Bits

	// PRR: Power Reduction Register
	PRR_PRLCD    = 0x10 // Power Reduction LCD
	PRR_PRTIM1   = 0x8  // Power Reduction Timer/Counter1
	PRR_PRSPI    = 0x4  // Power Reduction Serial Peripheral Interface
	PRR_PRUSART0 = 0x2  // Power Reduction USART
	PRR_PRADC    = 0x1  // Power Reduction ADC

	// SMCR: Sleep Mode Control Register
	SMCR_SM0 = 0x2 // Sleep Mode Select bits
	SMCR_SM1 = 0x4 // Sleep Mode Select bits
	SMCR_SM2 = 0x8 // Sleep Mode Select bits
	SMCR_SE  = 0x1 // Sleep Enable
)

// Bitfields for EXINT: External Interrupts
const (
	// EICRA: External Interrupt Control Register A
	EICRA_ISC01 = 0x2 // External Interrupt Sense Control 0 Bit 1
	EICRA_ISC00 = 0x1 // External Interrupt Sense Control 0 Bit 0

	// EIMSK: External Interrupt Mask Register
	EIMSK_PCIE0 = 0x10 // Pin Change Interrupt Enables
	EIMSK_PCIE1 = 0x20 // Pin Change Interrupt Enables
	EIMSK_PCIE2 = 0x40 // Pin Change Interrupt Enables
	EIMSK_PCIE3 = 0x80 // Pin Change Interrupt Enables
	EIMSK_INT0  = 0x1  // External Interrupt Request 0 Enable

	// EIFR: External Interrupt Flag Register
	EIFR_PCIF0 = 0x10 // Pin Change Interrupt Flags
	EIFR_PCIF1 = 0x20 // Pin Change Interrupt Flags
	EIFR_PCIF2 = 0x40 // Pin Change Interrupt Flags
	EIFR_PCIF3 = 0x80 // Pin Change Interrupt Flags
	EIFR_INTF0 = 0x1  // External Interrupt Flag 0
)

// Bitfields for EEPROM: EEPROM
const (
	// EECR: EEPROM Control Register
	EECR_EERIE = 0x8 // EEPROM Ready Interrupt Enable
	EECR_EEMWE = 0x4 // EEPROM Master Write Enable
	EECR_EEWE  = 0x2 // EEPROM Write Enable
	EECR_EERE  = 0x1 // EEPROM Read Enable
)

// Bitfields for TC8: Timer/Counter, 8-bit
const (
	// TCCR0A: Timer/Counter0 Control Register
	TCCR0A_FOC0A  = 0x80 // Force Output Compare
	TCCR0A_WGM00  = 0x40 // Waveform Generation Mode 0
	TCCR0A_COM0A0 = 0x10 // Compare Match Output Modes
	TCCR0A_COM0A1 = 0x20 // Compare Match Output Modes
	TCCR0A_WGM01  = 0x8  // Waveform Generation Mode 1
	TCCR0A_CS00   = 0x1  // Clock Selects
	TCCR0A_CS01   = 0x2  // Clock Selects
	TCCR0A_CS02   = 0x4  // Clock Selects

	// TIMSK0: Timer/Counter0 Interrupt Mask Register
	TIMSK0_OCIE0A = 0x2 // Timer/Counter0 Output Compare Match Interrupt Enable
	TIMSK0_TOIE0  = 0x1 // Timer/Counter0 Overflow Interrupt Enable

	// TIFR0: Timer/Counter0 Interrupt Flag register
	TIFR0_OCF0A = 0x2 // Timer/Counter0 Output Compare Flag 0
	TIFR0_TOV0  = 0x1 // Timer/Counter0 Overflow Flag
)

// Bitfields for TC8_ASYNC: Timer/Counter, 8-bit Async
const (
	// TCCR2A: Timer/Counter2 Control Register
	TCCR2A_FOC2A  = 0x80 // Force Output Compare A
	TCCR2A_WGM20  = 0x40 // Waveform Generation Mode
	TCCR2A_COM2A0 = 0x10 // Compare Output Mode bits
	TCCR2A_COM2A1 = 0x20 // Compare Output Mode bits
	TCCR2A_WGM21  = 0x8  // Waveform Generation Mode
	TCCR2A_CS20   = 0x1  // Clock Select bits
	TCCR2A_CS21   = 0x2  // Clock Select bits
	TCCR2A_CS22   = 0x4  // Clock Select bits

	// TIMSK2: Timer/Counter2 Interrupt Mask register
	TIMSK2_OCIE2A = 0x2 // Timer/Counter2 Output Compare Match Interrupt Enable
	TIMSK2_TOIE2  = 0x1 // Timer/Counter2 Overflow Interrupt Enable

	// TIFR2: Timer/Counter2 Interrupt Flag Register
	TIFR2_OCF2A = 0x2 // Timer/Counter2 Output Compare Flag 2
	TIFR2_TOV2  = 0x1 // Timer/Counter2 Overflow Flag

	// ASSR: Asynchronous Status Register
	ASSR_EXCLK  = 0x10 // Enable External Clock Interrupt
	ASSR_AS2    = 0x8  // AS2: Asynchronous Timer/Counter2
	ASSR_TCN2UB = 0x4  // TCN2UB: Timer/Counter2 Update Busy
	ASSR_OCR2UB = 0x2  // Output Compare Register2 Update Busy
	ASSR_TCR2UB = 0x1  // TCR2UB: Timer/Counter Control Register2 Update Busy
)

// Bitfields for TC16: Timer/Counter, 16-bit
const (
	// TCCR1A: Timer/Counter1 Control Register A
	TCCR1A_COM1A0 = 0x40 // Compare Output Mode 1A, bits
	TCCR1A_COM1A1 = 0x80 // Compare Output Mode 1A, bits
	TCCR1A_COM1B0 = 0x10 // Compare Output Mode 1B, bits
	TCCR1A_COM1B1 = 0x20 // Compare Output Mode 1B, bits
	TCCR1A_WGM10  = 0x1  // Waveform Generation Mode
	TCCR1A_WGM11  = 0x2  // Waveform Generation Mode

	// TCCR1B: Timer/Counter1 Control Register B
	TCCR1B_ICNC1 = 0x80 // Input Capture 1 Noise Canceler
	TCCR1B_ICES1 = 0x40 // Input Capture 1 Edge Select
	TCCR1B_WGM10 = 0x8  // Waveform Generation Mode
	TCCR1B_WGM11 = 0x10 // Waveform Generation Mode
	TCCR1B_CS10  = 0x1  // Prescaler source of Timer/Counter 1
	TCCR1B_CS11  = 0x2  // Prescaler source of Timer/Counter 1
	TCCR1B_CS12  = 0x4  // Prescaler source of Timer/Counter 1

	// TCCR1C: Timer/Counter 1 Control Register C
	TCCR1C_FOC1A = 0x80 // Force Output Compare 1A
	TCCR1C_FOC1B = 0x40 // Force Output Compare 1B

	// TIMSK1: Timer/Counter1 Interrupt Mask Register
	TIMSK1_ICIE1  = 0x20 // Timer/Counter1 Input Capture Interrupt Enable
	TIMSK1_OCIE1B = 0x4  // Timer/Counter1 Output Compare B Match Interrupt Enable
	TIMSK1_OCIE1A = 0x2  // Timer/Counter1 Output Compare A Match Interrupt Enable
	TIMSK1_TOIE1  = 0x1  // Timer/Counter1 Overflow Interrupt Enable

	// TIFR1: Timer/Counter1 Interrupt Flag register
	TIFR1_ICF1  = 0x20 // Input Capture Flag 1
	TIFR1_OCF1B = 0x4  // Output Compare Flag 1B
	TIFR1_OCF1A = 0x2  // Output Compare Flag 1A
	TIFR1_TOV1  = 0x1  // Timer/Counter1 Overflow Flag
)

// Bitfields for WDT: Watchdog Timer
const (
	// WDTCR: Watchdog Timer Control Register
	WDTCR_WDCE = 0x10 // Watchdog Change Enable
	WDTCR_WDE  = 0x8  // Watch Dog Enable
	WDTCR_WDP0 = 0x1  // Watch Dog Timer Prescaler bits
	WDTCR_WDP1 = 0x2  // Watch Dog Timer Prescaler bits
	WDTCR_WDP2 = 0x4  // Watch Dog Timer Prescaler bits
)

// Bitfields for BOOT_LOAD: Bootloader
const (
	// SPMCSR: Store Program Memory Control Register
	SPMCSR_SPMIE  = 0x80 // SPM Interrupt Enable
	SPMCSR_RWWSB  = 0x40 // Read While Write Section Busy
	SPMCSR_RWWSRE = 0x10 // Read While Write section read enable
	SPMCSR_BLBSET = 0x8  // Boot Lock Bit Set
	SPMCSR_PGWRT  = 0x4  // Page Write
	SPMCSR_PGERS  = 0x2  // Page Erase
	SPMCSR_SPMEN  = 0x1  // Store Program Memory Enable
)

// Bitfields for USART: USART
const (
	// UCSR0A: USART Control and Status Register A
	UCSR0A_RXC0  = 0x80 // USART Receive Complete
	UCSR0A_TXC0  = 0x40 // USART Transmit Complete
	UCSR0A_UDRE0 = 0x20 // USART Data Register Empty
	UCSR0A_FE0   = 0x10 // Framing Error
	UCSR0A_DOR0  = 0x8  // Data OverRun
	UCSR0A_UPE0  = 0x4  // USART Parity Error
	UCSR0A_U2X0  = 0x2  // Double the USART Transmission Speed
	UCSR0A_MPCM0 = 0x1  // Multi-processor Communication Mode

	// UCSR0B: USART Control and Status Register B
	UCSR0B_RXCIE0 = 0x80 // RX Complete Interrupt Enable
	UCSR0B_TXCIE0 = 0x40 // TX Complete Interrupt Enable
	UCSR0B_UDRIE0 = 0x20 // USART Data Register Empty Interrupt Enable
	UCSR0B_RXEN0  = 0x10 // Receiver Enable
	UCSR0B_TXEN0  = 0x8  // Transmitter Enable
	UCSR0B_UCSZ02 = 0x4  // Character Size
	UCSR0B_RXB80  = 0x2  // Receive Data Bit 8
	UCSR0B_TXB80  = 0x1  // Transmit Data Bit 8

	// UCSR0C: USART Control and Status Register C
	UCSR0C_UMSEL0 = 0x40 // USART Mode Select
	UCSR0C_UPM00  = 0x10 // Parity Mode Bits
	UCSR0C_UPM01  = 0x20 // Parity Mode Bits
	UCSR0C_USBS0  = 0x8  // Stop Bit Select
	UCSR0C_UCSZ00 = 0x2  // Character Size
	UCSR0C_UCSZ01 = 0x4  // Character Size
	UCSR0C_UCPOL0 = 0x1  // Clock Polarity
)
