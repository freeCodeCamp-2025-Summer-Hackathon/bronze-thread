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
import bag1 from '../assets/cta.png';
import image1 from '../assets/image1.jpg';
import image2 from '../assets/image2.jpg';
import image3 from '../assets/image3.jpg';
import image4 from '../assets/image4.jpg';
import image5 from '../assets/image5.jpg';

export const products: Product[] = [
  {
    id: 1,
    title: 'Canon AE-1 Film Camera',
    price: 280,
    image: image1,
    category: 'Electronics',
    thumbnails: [image2, image3, image4, image5],
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
    id: 2,
    title: 'Pentax K1000',
    price: 250,
    image: bag1,
    category: 'Electronics',
    thumbnails: [bag1, image2],
    listed: 'Listed 2 days ago',
    specs: ['35mm film format', 'FD lens mount', 'Manual focus'],
    seller: {
      name: 'John Doe',
      initials: 'JD',
      verified: false,
      location: 'Islamabad',
      rating: '4.8',
    },
    highlights: ['Classic build', 'Battery included'],
    description: 'Pentax K1000 is a durable and reliable 35mm film camera.',
  },
  {
    id: 3,
    title: 'Leica M6',
    price: 180,
    category: 'Electronics',
    image: bag1,
    thumbnails: [bag1, bag1],
    listed: 'Listed 2 days ago',
    specs: ['35mm film format', 'FD lens mount', 'Manual focus'],
    seller: {
      name: 'Anna Smith',
      initials: 'AS',
      verified: true,
      location: 'Islamabad',
      rating: '4.8',
    },
    highlights: ['Mint condition', 'With leather case'],
    description: 'Legendary Leica M6 rangefinder camera, perfect for collectors.',
  },
  {
    id: 4,
    title: 'Vintage Leather Handbag',
    price: 320,
    category: 'bags',
    image: bag1,
    thumbnails: [bag1, bag1],
    listed: 'Listed 2 days ago',
    specs: ['35mm film format', 'FD lens mount', 'Manual focus'],
    seller: {
      name: 'Liam Grant',
      initials: 'LG',
      verified: true,
      location: 'Islamabad',
      rating: '4.8',
    },
    highlights: ['Genuine leather', 'Spacious design'],
    description: 'Stylish vintage leather handbag with modern craftsmanship.',
  },
  {
    id: 5,
    title: 'Gold-Plated Bracelet',
    price: 150,
    category: 'accessories',
    image: bag1,
    listed: 'Listed 2 days ago',
    thumbnails: [bag1, bag1],
    specs: ['35mm film format', 'FD lens mount', 'Manual focus'],
    seller: {
      name: 'Emily Rose',
      initials: 'ER',
      verified: false,
      location: 'Islamabad',
      rating: '4.8',
    },
    highlights: ['Elegant design', 'Tarnish-resistant'],
    description: 'A beautiful gold-plated bracelet perfect for any outfit.',
  },
  {
    id: 6,
    title: 'Graphic Cotton T-Shirt',
    price: 200,
    category: 'clothing',
    listed: 'Listed 2 days ago',
    image: bag1,
    thumbnails: [bag1, bag1],
    specs: ['35mm film format', 'FD lens mount', 'Manual focus'],
    seller: {
      name: 'David Kim',
      initials: 'DK',
      verified: true,
      location: 'Islamabad',
      rating: '4.8',
    },
    highlights: ['100% cotton', 'Pre-shrunk fabric'],
    description: 'Comfortable cotton t-shirt with minimalist graphic print.',
  },
  {
    id: 7,
    title: 'Women’s Shoulder Tote',
    price: 310,
    listed: 'Listed 2 days ago',
    category: 'bags',
    image: bag1,
    thumbnails: [bag1, bag1],
    specs: ['35mm film format', 'FD lens mount', 'Manual focus'],
    seller: {
      name: 'Sophia Lee',
      initials: 'SL',
      verified: true,
      location: 'Islamabad',
      rating: '4.8',
    },
    highlights: ['Canvas material', 'Multiple compartments'],
    description: 'Trendy shoulder tote perfect for daily use or outings.',
  },
  {
    id: 8,
    title: 'Chunky Knit Sweater',
    price: 2100,
    category: 'clothing',
    listed: 'Listed 2 days ago',
    image: bag1,
    thumbnails: [bag1, bag1],
    specs: ['35mm film format', 'FD lens mount', 'Manual focus'],
    seller: {
      name: 'Robert Miles',
      initials: 'RM',
      verified: true,
      location: 'Islamabad',
      rating: '4.8',
    },
    highlights: ['Warm and cozy', 'Oversized fit'],
    description: 'Chunky knit sweater with premium yarn, great for cold seasons.',
  },
];
