word VentPin = 3;
int percentage = 25;
int mapedPercentage = 0;

void setup() {
  pinMode(VentPin, OUTPUT);
  Serial.begin(9600);
  pwm25kHzBegin();
}

void loop() {
  if (Serial.available()){
    String recievedChar = Serial.readStringUntil('\n');
    percentage = recievedChar.toInt();
  }
  int newPercentage = 0;
  if (percentage != 0) {
    newPercentage = (percentage - 1) * (0 - 79) / (100 - 1) + 79;
  } else {
    newPercentage = 0;
  }

  if (newPercentage != mapedPercentage) {
    Serial.println(percentage);
    pwmDuty(newPercentage); // 75% (range = 0-79 = 1.25-100%)
    mapedPercentage = newPercentage;
  }
}

void pwm25kHzBegin() {
  TCCR2A = 0;                               // TC2 Control Register A
  TCCR2B = 0;                               // TC2 Control Register B
  TIMSK2 = 0;                               // TC2 Interrupt Mask Register
  TIFR2 = 0;                                // TC2 Interrupt Flag Register
  TCCR2A |= (1 << COM2B1) | (1 << WGM21) | (1 << WGM20);  // OC2B cleared/set on match when up/down counting, fast PWM
  TCCR2B |= (1 << WGM22) | (1 << CS21);     // prescaler 8
  OCR2A = 79;                               // TOP overflow value (Hz)
  OCR2B = 0;
}

void pwmDuty(byte ocrb) {
  OCR2B = ocrb;                             // PWM Width (duty)
}
