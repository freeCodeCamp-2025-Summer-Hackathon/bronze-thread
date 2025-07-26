export interface Product {
  id: number;
  title: string;
  image: string;
  price: number;
  category: string;
  thumbnails: string[];
  specs: string[];
  listed: string;
  seller: {
    name: string;
    initials: string;
    verified: boolean;
    location: string;
    rating: string;
  };

  highlights: string[];
  description: string;
}
// Camera Images
import canon1 from "@/assets/CanonAE-1a.jpg";
import canonThumb1 from "@/assets/CanonAE-1b.jpg";
import canonThumb2 from "@/assets/CanonAE-1c.jpg";
import canonThumb3 from "@/assets/CanonAE-1d.jpg";
import canonThumb4 from "@/assets/CanonAE-1e.jpg";

import pentax1 from "@/assets/PentaxK1000a.jpg";
import pentaxThumb1 from "@/assets/PentaxK1000b.jpg";
import pentaxThumb2 from "@/assets/PentaxK1000ac.jpg";
import pentaxThumb3 from "@/assets/PentaxK1000d.jpg";

import leica1 from "@/assets/LeicaM6a.jpg";
import leicaThumb1 from "@/assets/LeicaM6b.jpg";
import leicaThumb2 from "@/assets/LeicaM6c.jpg";
import leicaThumb3 from "@/assets/LeicaM6d.jpg";

import leatherBag1 from "@/assets/LeatherTote1.jpg";
import leatherBag2 from "@/assets/LeatherTote2.jpg";
import leatherBag3 from "@/assets/LeatherTote3.jpg";
import leatherBag4 from "@/assets/LeatherTote4.jpg";

import tote1 from "@/assets/tote1.jpg";
import tote2 from "@/assets/tote2.jpg";
import tote3 from "@/assets/tote3.jpg";
import tote4 from "@/assets/tote4.jpg";

import miniPurse1 from "@/assets/minipurse1.jpg";
import miniPurse2 from "@/assets/minipurse2.jpg";
import miniPurse3 from "@/assets/minipurse3.jpg";
import miniPurse4 from "@/assets/minipurse4.jpg";

import backpack1 from "@/assets/backpack1.jpg";
import backpack2 from "@/assets/backpack2.jpg";
import backpack3 from "@/assets/backpack3.jpg";
import backpack4 from "@/assets/backpack4.jpg";

import officeBag1 from "@/assets/officeBag1.jpg";
import officeBag2 from "@/assets/officeBag2.jpg";
import officeBag3 from "@/assets/officeBag3.jpg";
import officeBag4 from "@/assets/officeBag4.jpg";

import bracelet1 from "@/assets/Bracelet1.jpg";
import bracelet2 from "@/assets/Bracelet2.jpg";
import bracelet3 from "@/assets/Bracelet3.jpg";
import bracelet4 from "@/assets/Bracelet4.jpg";

import clutch1 from "@/assets/clutch1.jpg";
import clutch2 from "@/assets/clutch2.jpg";
import clutch3 from "@/assets/clutch3.jpg";
import clutch4 from "@/assets/clutch4.jpg";

import tshirt1 from "@/assets/Tshirt1.jpg";
import tshirt2 from "@/assets/Tshirt2.jpg";
import tshirt3 from "@/assets/Tshirt3.jpg";
import tshirt4 from "@/assets/Tshirt4.jpg";

import sweater1 from "@/assets/sweater1.jpg";
import sweater2 from "@/assets/sweater2.jpg";
import sweater3 from "@/assets/sweater3.jpg";
import sweater4 from "@/assets/sweater4.jpg";


export const products: Product[] = [
  {
    id: 1,
    title: 'Graphic Cotton T-Shirt',
    price: 200,
    category: 'Clothing',
    listed: 'Listed 2 days ago',
    image: tshirt1,
    thumbnails: [
      tshirt2,
      tshirt3,
      tshirt4
    ],
    specs: [
      'Fabric: 100% ring-spun cotton',
      'Pre-shrunk for consistent fit',
      'Short sleeves with ribbed collar',
      'Machine washable',
      'Printed with eco-friendly inks',
      'Unisex fit'
    ],
    seller: {
      name: 'David Kim',
      initials: 'DK',
      verified: true,
      location: 'Islamabad',
      rating: '4.8'
    },
    highlights: ["Tote Bag", "Clutch", "Aesthetic Hair Clips"],
    description:
      'Express yourself in comfort with this graphic cotton t-shirt. Its minimalist design and breathable fabric make it a staple piece for casual wardrobes. Designed for all-day wear and lasting style.'
  },

  
  {
    id: 2,
    title: 'Gold-Plated Bracelet',
    price: 150,
    category: 'Accessories',
    image: bracelet1,
    listed: 'Listed 2 days ago',
    thumbnails: [
      bracelet2,
      bracelet3,
      bracelet4
    ],
    specs: [
      'Material: Stainless steel with gold plating',
      'Adjustable chain length (16-20cm)',
      'Tarnish-resistant coating',
      'Hypoallergenic finish',
      'Secure lobster clasp'
    ],
    seller: {
      name: 'Emily Rose',
      initials: 'ER',
      verified: false,
      location: 'Islamabad',
      rating: '4.8'
    },
    highlights: ["Sketchbook", "Resin Rings", "Phone Charm"],
    description:
      'This delicate gold-plated bracelet is ideal for everyday elegance. Whether worn solo or stacked, it adds a graceful shine to your look. Lightweight, minimal, and ideal as a gift or personal accessory.'
  },

  {
    id: 3,
    title: "Mini Leather Purse",
    price: 170,
    category: "Bags",
    image: miniPurse1,
    listed: "Listed 2 days ago",
    thumbnails: [
      miniPurse2,
      miniPurse3,
      miniPurse4
    ],
    specs: [
      "Material: Premium PU Leather",
      "Compact and lightweight design",
      "Secure magnetic snap closure",
      "Gold-tone hardware",
      "Single interior compartment",
      "Adjustable crossbody strap"
    ],
    seller: {
      name: "Nadia Malik",
      initials: "NM",
      verified: true,
      location: "Lahore",
      rating: "4.7"
    },
     highlights: ["Laptop Sleeve", "Shoes", "Sling Bag"],
    description:
      "Chic and minimal, this Mini Leather Purse is a perfect blend of elegance and utility. Ideal for casual hangouts and formal gatherings, it holds your essentials with style."
  },

  {
    id: 4,
    title: 'Vintage Leather Handbag',
    price: 320,
    category: 'Bags',
    image: leatherBag1,
    listed: 'Listed 2 days ago',
    thumbnails: [
      leatherBag2,
      leatherBag3,
      leatherBag4
    ],
    specs: [
      'Material: 100% genuine leather',
      'Magnetic flap + zipper closure',
      'Spacious inner compartments',
      'Inner lining with zip pocket',
      'Dual top handles',
      'Handcrafted design'
    ],
    seller: {
      name: 'Liam Grant',
      initials: 'LG',
      verified: true,
      location: 'Islamabad',
      rating: '4.8'
    },
    highlights: ["Mini Purse", "Jewelry", "Branded Sunglasses"],
    description:
      'This vintage leather handbag is both rustic and refined. Crafted from full-grain leather, it’s ideal for daily use while giving a nod to timeless style. The spacious interior makes it perfect for work, shopping, or a casual day out.'
  },
  {
    id: 5,
    title: 'Pentax K1000',
    price: 250,
    image: pentax1,
    category: 'Electronics',
    listed: 'Listed 2 days ago',
    thumbnails: [
      pentaxThumb1,
      pentaxThumb2,
      pentaxThumb3
    ],
    specs: [
      '35mm Single-Lens Reflex (SLR) Camera',
      'K-mount for interchangeable lenses',
      'Fully manual controls for shutter speed and aperture',
      'Shutter Speed: 1/1000 to 1 second + Bulb mode',
      'Center-weighted metering system',
      'Built-in light meter powered by 1x LR44 battery',
      'Weight: approx. 620g',
      'Dimensions: 143 x 91 x 49 mm'
    ],
    seller: {
      name: 'John Doe',
      initials: 'JD',
      verified: false,
      location: 'Islamabad',
      rating: '4.8'
    },
    highlights: ["Mini Tripod", "Camera Strap", "Film Rolls"],
    description:
      'The Pentax K1000 is a legendary 35mm film camera known for its durability and simplicity. Widely used by photography schools, it offers manual control and reliable performance, making it perfect for both learners and enthusiasts looking for full creative freedom.'
  },
 {
    id: 6,
    title: "Urban Explorer Backpack",
    price: 450,
    category: "Bags",
    image: backpack1,
    listed: "Listed 2 days ago",
    thumbnails: [
      backpack2,
      backpack3,
      backpack4
    ],
    specs: [
      "Material: Durable canvas with faux leather straps",
      "Laptop compartment fits up to 15.6 inches",
      "Water-resistant lining",
      "Multiple organizer pockets",
      "Zipper + magnetic flap closure",
      "Capacity: 20L"
    ],
    seller: {
      name: "Zain Rauf",
      initials: "ZR",
      verified: true,
      location: "Karachi",
      rating: "4.9"
    },
    highlights: ["Gold Earrings", "Pearl Bracelet", "Mini Shoulder Bag"],
    description:
      "Made for modern explorers, this backpack offers a rugged look with practical storage. Whether you're heading to class or a day trip, it's your go-to carry-all companion."
  },

  {
    id: 7,
    title: 'Women’s Shoulder Tote',
    price: 310,
    listed: 'Listed 2 days ago',
    category: 'Bags',
    image: tote1,
    thumbnails: [
      tote2,
      tote3,
      tote4
    ],
    specs: [
      'Material: Canvas & PU leather blend',
      'Main zipper closure',
      'Interior zip and slip pockets',
      'Double-stitched handles',
      'Reinforced bottom panel',
      'Fits A4 files and 13-inch laptop'
    ],
    seller: {
      name: 'Sophia Lee',
      initials: 'SL',
      verified: true,
      location: 'Islamabad',
      rating: '4.8'
    },
    highlights: ["Heels", "Blazer", "Card Holder"],
    description:
      'A daily essential with a modern edge, this shoulder tote bag fits everything from laptops to groceries. The minimalist design and durable build make it your go-to carryall.'
  },

  {
    id: 8,
    title: 'Chunky Knit Sweater',
    price: 2100,
    category: 'Clothing',
    listed: 'Listed 2 days ago',
    image: sweater1,
    thumbnails: [
      sweater2,
      sweater3,
      sweater4
    ],
    specs: [
      'Material: Wool-acrylic blend',
      'Loose oversized fit',
      'Ribbed hem and cuffs',
      'Crew neck design',
      'Machine washable (cold)',
      'Soft brushed texture'
    ],
    seller: {
      name: 'Robert Miles',
      initials: 'RM',
      verified: true,
      location: 'Islamabad',
      rating: '4.8'
    },
     highlights: ["Film Packs", "Cute Backpack", "Travel Journal"],
    description:
      'Wrap yourself in warmth with this chunky knit sweater. Its relaxed silhouette and soft finish provide ultimate comfort without sacrificing style — ideal for chilly evenings and lazy Sundays.'
  },
  {
    id: 9,
    title: 'Leica M6',
    price: 180,
    category: 'Electronics',
    image: leica1,
    listed: 'Listed 2 days ago',
    thumbnails: [
      leicaThumb1,
      leicaThumb2,
      leicaThumb3
    ],
    specs: [
      '35mm Rangefinder Film Camera',
      'Leica M-mount lenses compatible',
      'Manual focus with bright viewfinder',
      'Shutter Speeds: 1/1000 to 1 sec + Bulb mode',
      'Through-the-lens light metering',
      '2 LR44 batteries required for meter',
      'Compact & lightweight build: 560g',
      'Dimensions: 138 x 77 x 38 mm'
    ],
    seller: {
      name: 'Anna Smith',
      initials: 'AS',
      verified: true,
      location: 'Islamabad',
      rating: '4.8'
    },
   highlights: ["High-end DSLR", "Luxury Watch", "Vintage Leather Jacket"],
    description:
      'The Leica M6 is one of the most revered rangefinder film cameras of all time. With its elegant mechanical build and precise manual operation, it’s a favorite among collectors and street photographers alike. Comes with original accessories and immaculate optics.'
  },
{
    id: 10,
    title: 'Canon AE-1 Film Camera',
    price: 280,
    image: canon1,
    category: 'Electronics',
    thumbnails: [canonThumb1, canonThumb2, canonThumb3, canonThumb4],
    listed: 'Listed 2 days ago',
    specs: [
      '35mm Single-Lens Reflex (SLR) Film Camera',
      ' 135 (35mm standard film)',
      ' Canon FD Mount (interchangeable lenses supported)',
      'Shutter Speed Range is 2 seconds – 1/1000 sec + Bulb (B) mode',
      'Light Metering is Center-weighted TTL (Through-the-Lens) metering',
      ' ISO 25 – 3200 (manually set)',
      'Battery Requirement is 1x 4LR44 or PX28 6V battery (required for all operations)',
      'Weight (Body Only) is Approx. 590g (1.3 lbs)',
      'Dimensions 141 x 87 x 47.5 mm',
    ],
    seller: {
      name: 'Sarah Chaudry',
      initials: 'SC',
      verified: true,
      location: 'Islamabad',
      rating: '4.8',
    },
    highlights: [
      'Vintage Vinyl Records (Jazz/Blues)',
      'Film Photography Books',
      'Darkroom Equipment',
      'Other Vintage Cameras',
      'Art Supplies (Professional Grade)',
    ],
    description:
      'Capture timeless memories with the iconic Canon AE-1, a 35mm film camera loved by professionals and hobbyists alike.Its classic design, manual controls, and sharp FD lens compatibility make every shot intentional and beautiful.Perfect for film photography enthusiasts seeking rich, nostalgic tones and full creative control.Whether you are starting your analog journey or expanding your vintage collection — this camera delivers.',
  },
  
 
  {
    id: 11,
    title: "Evening Satin Clutch",
    price: 220,
    category: "Accessories",
    image: clutch1,
    listed: "Listed 2 days ago",
    thumbnails: [
      clutch2,
      clutch3,
      clutch4
    ],
    specs: [
      "Material: Soft satin exterior",
      "Magnetic button closure",
      "Detachable metal chain strap",
      "Interior slip pocket",
      "Fits phone, keys, and cards"
    ],
    seller: {
      name: "Ayesha Noor",
      initials: "AN",
      verified: false,
      location: "Islamabad",
      rating: "4.6"
    },
     highlights: ["Cargos", "Oversized Hoodie", "Denim Jacket"],
    description:
      "This satin clutch adds a glamorous touch to your outfit. With a sleek structure and soft finish, it’s perfect for weddings, formal dinners, and date nights."
  },
  {
    id: 12,
    title: "Professional Office Bag",
    price: 650,
    category: "Bags",
    image: officeBag1,
    listed: "Listed 2 days ago",
    thumbnails: [
      officeBag2,
      officeBag3,
      officeBag4
    ],
    specs: [
      "Material: Vegan leather",
      "Padded laptop compartment",
      "Zipper and magnetic closure",
      "Multiple internal organizers",
      "Reinforced handles + detachable shoulder strap",
      "Fits documents, chargers, and tech accessories"
    ],
    seller: {
      name: "Hamza Tariq",
      initials: "HT",
      verified: true,
      location: "Islamabad",
      rating: "5.0"
    },
    highlights: ["Crochet Top", "Phone Strap", "Mini Hair Claw Set"],
    description:
      "Crafted for professionals, this Office Bag keeps you organized in style. With ample space and refined design, it's the ideal companion for meetings and daily commute."
  },

];
