import sys
import json

muni_filename = "mo/stl_munis.json"
court_filename = "mo/stl_courts.json"

def point_is_in_bounds(x, y, bounds):
    #Checks if the point (x, y) is contained in the region bounded by the bounds variable
    #Code based on algorithm found at http://www.ariel.com.au/a/python-point-int-poly.html
    n = len(bounds)
    inside = False

    p1x,p1y = bounds[0]
    for i in range(n+1):
        p2x,p2y = bounds[i % n]
        if y > min(p1y,p2y):
            if y <= max(p1y,p2y):
                if x <= max(p1x,p2x):
                    if p1y != p2y:
                        xinters = (y-p1y)*(p2x-p1x)/(p2y-p1y)+p1x
                    if p1x == p2x or x <= xinters:
                        inside = not inside
        p1x,p1y = p2x,p2y

    return inside

def point_is_in_poly(x, y, poly):
    #The first linearRing is defined in the geojson standard to always be the polygon's outer boundary.
    #If a point is in the polygon, it must be inside this boundary and outside all subsequent boundaries.
    outer_bound = poly[0]
    
    if point_is_in_bounds(x, y, outer_bound):
        if len(poly) == 1:
            #Polygon has no holes and the point is inside its outer boundary. Polygon contains the point.
            return True
        else:
            #Polygon has at least one hole. Check to make sure the point isn't in any of the polygon's holes.
            for hole in poly[1:len(poly)]:
                if point_is_in_bounds(x, y, hole):
                    #If in a hole, it's not in the polygon. Return false immediately.
                    return False
            #If it's not in any of the holes but is inside the outermost bound, it's in the polygon. Return true.
            return True
    else:
        #Point is outside of the polygon's outer bounding ring.
        return False

def point_is_in_multipoly(x, y, multipoly):
    #Determines if (x, y) is contaiend in the multipolygon multipoly
    print("MultiPoly of ", len(multipoly), "polygons.")
    for poly in multipoly:
        print("Next poly has", len(poly), "points.")
        if point_is_in_poly(x, y, poly):
            return True
    return False

def muniForPoint(x, y):
    #Load municipal boundary data into JSON objects readable by Python
    file = open(muni_filename, 'r')
    muni_json_data = json.loads(file.read())
    file.close()
    munis = muni_json_data["features"]

    #Iterate through those objects and look for one containing the point (x, y) within its bounds.
    for muni in munis:
        #Grab the geometry so we know what type of polygon we're dealing with.
        muni_geometry = muni["geometry"]
        print("Checking for match in:", muni["properties"]["MUNICIPALITY"])
        if muni_geometry["type"] == "MultiPolygon":
            #Geometry is a multipolygon, iterate through all of its polygons
            multipoly = muni_geometry["coordinates"]
            if point_is_in_multipoly(x, y, multipoly):
                #Return if we have a match.
                return muni
        elif muni_geometry["type"] == "Polygon":
            #Geometry is a regular polygon, use normal method
            poly = muni_geometry["coordinates"]
            if point_is_in_poly(x, y, poly):
                #Return if we have a match.
                return muni
        else:
            print("ERROR. INVALID GEOMETRY TYPE")
    #Return nothing if there's no match.
    return None

def courtForMuni(muni):
    file = open(court_filename, 'r')
    court_json_data = json.loads(file.read())
    file.close()
    courts = court_json_data["features"]

    for court in courts:
        if court["properties"]["Municipali"].lower() == muni["properties"]["MUNICIPALITY"].lower():
            return court
## Test
def main():
    #print("x coordinate?")
    x = sys.argv[0]#eval(input())
    #print("y coordinate?")
    y = sys.argv[1]#eval(input())
    muni = muniForPoint(x, y)
    court = courtForMuni(muni)
    if muni == None:
        print("No municipality was found for that point.")
    else:
        print(json.dumps(muni["properties"]) + " :^) " + json.dumps(court))
main()

