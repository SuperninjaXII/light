
let clicked = (elem) => {
  let tl = gsap.timeline()
  tl.to(".options div",{
    width:'9svw',
    duration:.2,
    background:"transparent"
  })
  tl.to(elem, {
    width: '13svw',
    duration: 0.2,
    background:"#D39D46"
  });
  tl.to(".options div img",{
    width:'1.2rem',
    background:'transparent'
    })
};

let containers = document.querySelectorAll('.options div');

// Add event listeners to each div
containers.forEach(container => {
  container.addEventListener("click", function(event) {
    event.stopPropagation(); 
    clicked(event.target);
  });
});

document.querySelector('body').addEventListener("click", () => {
  gsap.to(".options div", {
    width: "20svw",
    duration: 0.2,
    background:"transparent",
  });
  gsap.to(".options div img",{
    width:'1rem',
    })
});