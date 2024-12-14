// Hero Slider
const slides = document.querySelectorAll('.slide');
const dots = document.querySelector('.slide-dots');
const prevBtn = document.querySelector('.prev-slide');
const nextBtn = document.querySelector('.next-slide');
let currentSlide = 0;

// Create dots for each slide
slides.forEach((_, index) => {
    const dot = document.createElement('div');
    dot.classList.add('dot');
    if (index === 0) dot.classList.add('active');
    dot.addEventListener('click', () => goToSlide(index));
    dots.appendChild(dot);
});

// Function to update slides
function updateSlides() {
    slides.forEach((slide, index) => {
        if (index === currentSlide) {
            slide.style.transform = 'translateX(0)';
            slide.style.opacity = '1';
        } else if (index < currentSlide) {
            slide.style.transform = 'translateX(-100%)';
            slide.style.opacity = '0';
        } else {
            slide.style.transform = 'translateX(100%)';
            slide.style.opacity = '0';
        }
    });

    // Update dots
    document.querySelectorAll('.dot').forEach((dot, index) => {
        dot.classList.toggle('active', index === currentSlide);
    });
}

// Next slide function
function nextSlide() {
    currentSlide = (currentSlide + 1) % slides.length;
    updateSlides();
}

// Previous slide function
function prevSlide() {
    currentSlide = (currentSlide - 1 + slides.length) % slides.length;
    updateSlides();
}

// Go to specific slide
function goToSlide(index) {
    currentSlide = index;
    updateSlides();
}

// Event listeners for controls
prevBtn.addEventListener('click', () => {
    prevSlide();
    resetTimer();
});

nextBtn.addEventListener('click', () => {
    nextSlide();
    resetTimer();
});

// Auto slide functionality
let slideInterval = null;

function startTimer() {
    slideInterval = setInterval(nextSlide, 3000); // Change slide every 3 seconds
}

function resetTimer() {
    if (slideInterval) {
        clearInterval(slideInterval);
        startTimer();
    }
}

// Start auto-sliding
startTimer();

// Pause on hover
const heroSection = document.querySelector('.hero');
heroSection.addEventListener('mouseenter', () => {
    clearInterval(slideInterval);
    slideInterval = null;
});

heroSection.addEventListener('mouseleave', () => {
    startTimer();
});

// Initialize first slide
updateSlides();